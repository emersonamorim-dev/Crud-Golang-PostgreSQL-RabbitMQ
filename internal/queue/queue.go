package queue

import (
	"log"

	"github.com/streadway/amqp"
)

// representa a conexão com o RabbitMQ.
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Define os métodos para interagir com a fila
type QueueClient interface {
	Publish(queueName string, msg []byte) error
}

// Garante que RabbitMQ implemente a interface QueueClient
var _ QueueClient = (*RabbitMQ)(nil)

// NewRabbitMQ inicializa e retorna uma nova instância do RabbitMQ.
func NewRabbitMQ(conn *amqp.Connection) (*RabbitMQ, error) {
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

// Publish publica uma mensagem na fila especificada.
func (r *RabbitMQ) Publish(queueName string, msg []byte) error {
	// Garante que a fila exista antes de publicar
	_, err := r.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Config a mensagem para publicação
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}

	// Publica a mensagem na fila
	err = r.channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,
		false,
		message,
	)
	return err
}

// Encerra a conexão e o canal com o RabbitMQ.
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			log.Printf("Erro ao fechar o canal do RabbitMQ: %v", err)
		}
	}
	if r.conn != nil {
		if err := r.conn.Close(); err != nil {
			log.Printf("Erro ao fechar a conexão do RabbitMQ: %v", err)
		}
	}
}
