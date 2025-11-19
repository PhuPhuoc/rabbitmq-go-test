package main

import (
	"github.com/PhuPhuoc/rabbitmq-go-test/publisher/config"
	httpserver "github.com/PhuPhuoc/rabbitmq-go-test/publisher/http"
	"github.com/PhuPhuoc/rabbitmq-go-test/publisher/queue"
)

func main() {
	configs := config.LoadConfig()

	// Connect to RabbitMQ
	conn := queue.InitRabbitMQ(configs)
	defer conn.Close()

	// Start HTTP server
	server := httpserver.NewHTTPServer(conn)
	server.Start(configs.APP_PORT)
}
