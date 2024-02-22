package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Crud-Golang-RabbitMQ/internal/service"
)

func main() {
	// Config da conexão com o RabbitMQ
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("Falha ao conectar-se ao RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Falha ao abrir um canal: %v", err)
	}

	// Config da conexão com PostgreSQL
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Não é possível se conectar ao banco de dados: %v", err)
	}

	// Inicialização do serviço de consumidor
	queueName := os.Getenv("RABBITMQ_QUEUE_NAME")
	consumerService := service.NewRabbitMQConsumer(channel, db, queueName)
	consumerService.StartConsumer()

	// O consumidor está rodando e processando mensagens
	log.Println("Consumidor iniciado e aguardando mensagens...")
	select {} 
}
