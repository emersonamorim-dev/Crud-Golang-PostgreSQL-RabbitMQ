package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	// err := godotenv.Load("/mnt/wsl/Crud-Golang-RabbitMQ/tests/.env")
	err := godotenv.Load("/home/seu-usuario/Projetos-Golang/Crud-Golang-RabbitMQ/tests/.env")

	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env:", err)
	}

	// Busca informações de conexão do ambiente
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Variável de ambiente DATABASE_URL não definida")
	}

	// Abre uma conexão com o BD PostgreSQL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Tenta pingar o BD para verificar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados PostgreSQL estabelecida com sucesso!")

	// Cria uma conexão AMQP usando as variáveis de ambiente
	rabbitMQPortStr := os.Getenv("RABBITMQ_PORT")
	rabbitMQPort, err := strconv.Atoi(rabbitMQPortStr)
	if err != nil {
		log.Fatalf("Erro ao converter a porta do RabbitMQ para int: %v", err)
	}

	// Testa a conexão com o RabbitMQ
	rabbitMQConnectionString := fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("RABBITMQ_USER"),
		os.Getenv("RABBITMQ_PASSWORD"),
		os.Getenv("RABBITMQ_HOST"),
		os.Getenv("RABBITMQ_PORT"))

	connRabbitMQ, err := amqp.Dial(rabbitMQConnectionString)
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	defer connRabbitMQ.Close()

	log.Println("Conexão com o RabbitMQ estabelecida com sucesso!")

	connectionString := amqp.URI{
		Scheme:   "amqp",
		Host:     os.Getenv("RABBITMQ_HOST"),
		Port:     rabbitMQPort,
		Username: os.Getenv("RABBITMQ_USER"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	}.String()

	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	defer conn.Close()

	log.Println("Conexão estabelecida com sucesso!")
}
