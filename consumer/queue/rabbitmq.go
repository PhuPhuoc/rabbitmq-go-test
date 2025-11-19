package queue

import (
	"fmt"

	"github.com/PhuPhuoc/rabbitmq-go-test/consumer/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
}

func InitRabbitMQ(config *config.Config) *RabbitConfig {
	var qconfig = RabbitConfig{}
	qconfig.connectRabbitMQ(config.RABBIT_USERNAME, config.RABBIT_PASSWORD, config.RABBIT_HOST, config.RABBIT_PORT)
	qconfig.setupChannelAndQueue()
	return &qconfig
}

func (q *RabbitConfig) connectRabbitMQ(username, pwd, host, port string) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", username, pwd, host, port))
	if err != nil {
		logrus.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	q.conn = conn
}

func (q *RabbitConfig) setupChannelAndQueue() error {
	// Create channel
	ch, err := q.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	q.Channel = ch

	// Setup queue
	queue, err := ch.QueueDeclare(
		"publisher",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}
	logrus.Infof("Queue declared: %s", queue.Name)
	q.QueueName = queue.Name

	return nil
}

func (r *RabbitConfig) Close() {
	if r.Channel != nil {
		_ = r.Channel.Close()
	}
	if r.conn != nil {
		_ = r.conn.Close()
	}
}
