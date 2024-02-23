package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Crud-Golang-RabbitMQ/internal/handler"
	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/queue"
	"Crud-Golang-RabbitMQ/internal/repository"
	"Crud-Golang-RabbitMQ/internal/routes"
	"Crud-Golang-RabbitMQ/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Printf("Aviso: Arquivo .env não encontrado. Continuando com variáveis de ambiente do sistema.")
	}

	// Define a URL de conexão com o RabbitMQ a partir de uma variável de ambiente
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		log.Fatal("A variável RABBITMQ_URL não está definida")
	}

	// Config do servidor Gin
	r := gin.Default()

	// Config e conexão com o PostgreSQL
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Falha ao inicializar o banco de dados: %v", err)
	}

	// Config e conexão com o RabbitMQ
	rabbitMQConn, err := setupRabbitMQ()
	if err != nil {
		log.Fatalf("Falha ao conectar-se a RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	// Cria um novo objeto RabbitMQ usando o wrapper definido no pacote `queue`
	rabbitMQClient, err := queue.NewRabbitMQ(rabbitMQConn)
	if err != nil {
		log.Fatalf("Falha ao criar o objeto RabbitMQ: %v", err)
	}

	// Inicializa repository, service, e handler
	clientesRepo := repository.NewClientesRepository(db)
	clienteService := service.NewClienteService(clientesRepo, rabbitMQClient)
	clientesHandler := handler.NewClientesHandler(clienteService, rabbitMQClient)

	// Config das rotas
	routes.SetupRouter(r, clientesHandler)

	// Inicia o servidor Gin em uma goroutine para permitir graceful shutdown
	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Falha ao iniciar servidor: %v", err)
		}
	}()

	// Aguarda sinais de interrupção graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Desligando servidor...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Forçando desligamento do servidor:", err)
	}

	log.Println("Servidor desligado com sucesso.")
}

func setupDatabase() (*gorm.DB, error) {
	postgresURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao banco de dados: %v", err)
	}

	// Executa migrações
	err = db.AutoMigrate(&model.Clientes{})
	if err != nil {
		return nil, fmt.Errorf("falha ao migrar: %v", err)
	}

	return db, nil
}

func setupRabbitMQ() (*amqp.Connection, error) {
	rabbitMQURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"),
	)

	return amqp.Dial(rabbitMQURL)
}
