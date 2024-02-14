package main

import (
	"Crud-Golang-RabbitMQ/internal/model"
	"Crud-Golang-RabbitMQ/internal/repository"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Carrega variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro a carregar .env file: %v", err)
	}

	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	postgresURL := os.Getenv("POSTGRES_URL")

	// Conecta ao PostgreSQL
	db, err := connectToGormDB(postgresURL)
	if err != nil {
		log.Fatalf("Não é possível se conectar ao BD: %v", err)
	}
	// Obtem a conexão SQL DB subjacente e fecha
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Falha ao obter BD SQL de gorm DB: %v", err)
	}
	defer sqlDB.Close()

	// Conecta ao RabbitMQ
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Falha ao conectar-se a RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir um canal: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"clientesQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Não foi possível declarar uma queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Falha ao cadastrar um consumidor: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recebeu uma mensagem: %s", d.Body)

			// Processa a mensagem
			if err := processMessage(db, d.Body); err != nil {
				log.Printf("Erro ao processar mensagem: %s", err)
			}
		}
	}()

	log.Printf(" [*] Aguardando mensagens.")
	<-forever
}

// Inicializa o repositório de Clientes com o pool de conexões com BD
func processMessage(db *gorm.DB, messageBody []byte) error {
	var cliente model.Clientes
	err := json.Unmarshal(messageBody, &cliente)
	if err != nil {
		log.Printf("Erro ao desempacotar mensagem: %v", err)
		return err
	}

	repo := repository.NewClientesRepository(db)

	// Método para criar um novo cliente no BD
	err = repo.Create(context.Background(), cliente)
	if err != nil {
		log.Printf("Erro ao inserir cliente no banco de dados: %v", err)
		return err
	}

	log.Printf("Cliente processado com sucesso: %+v", cliente)
	return nil
}

// Conexão do PostgreSQL usando Gorm
func connectToGormDB(postgresURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(postgresURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
