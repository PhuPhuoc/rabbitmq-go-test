package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PhuPhuoc/rabbitmq-go-test/publisher/queue"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	rabbitmq *queue.RabbitConfig
}

func NewHTTPServer(rabbitmq *queue.RabbitConfig) *HTTPServer {
	return &HTTPServer{
		rabbitmq: rabbitmq,
	}
}

func (s *HTTPServer) Start(port string) {
	http.HandleFunc("/publish", s.handlePublish)

	logrus.Infof("HTTP server running on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type PublishRequest struct {
	Message string `json:"message"`
}

func (s *HTTPServer) handlePublish(w http.ResponseWriter, r *http.Request) {
	// Chỉ cho phép POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PublishRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Message is required", http.StatusBadRequest)
		return
	}

	logrus.Infof("Received message: %s", req.Message)

	err = PublishMessage(s.rabbitmq, req.Message)
	if err != nil {
		logrus.Errorf("Failed to publish message: %v", err)
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("publish success!"))
}
