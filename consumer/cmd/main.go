package main

import (
	"github.com/PhuPhuoc/rabbitmq-go-test/consumer/config"
	"github.com/PhuPhuoc/rabbitmq-go-test/consumer/queue"
)

func main() {
	configs := config.LoadConfig()

	// Connect to RabbitMQ
	conn := queue.InitRabbitMQ(configs)
	defer conn.Close()

	// Start HTTP server
	// server := NewHTTPServer.NewHTTPServer(conn)
	// server.Start(configs.APP_PORT)
}
