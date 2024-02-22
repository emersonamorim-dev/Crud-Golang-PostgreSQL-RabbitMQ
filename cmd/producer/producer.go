package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

// publishMessage envia uma msg para a fila do RabbitMQ.
func publishMessage(channel *amqp.Channel, queueName, message string) error {
	err := channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("mensagem de erro ao publicar: %s", err)
	}
	return nil
}
