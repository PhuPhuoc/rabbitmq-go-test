package httpserver

import (
	"fmt"

	"github.com/PhuPhuoc/rabbitmq-go-test/publisher/queue"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func PublishMessage(q *queue.RabbitConfig, message string) error {
	err := q.Channel.Publish(
		"",          // exchange
		q.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %w", err)
	}

	logrus.Println("Publish success!")
	return nil
}
