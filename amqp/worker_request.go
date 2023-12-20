package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerRequest struct {
	RequestId   string   `json:"request_id"`
	RequestType string   `json:"request_type"`
	Angle       int64    `json:"angle,omitempty"`
	ImageType   string   `json:"image_type,omitempty"`
	PageNumbers []int    `json:"page_numbers,omitempty"`
	Ranges      []string `json:"ranges,omitempty"`
}

var creatworkerreqq sync.Once

func (a *RabbitAMQPClient) SendWorkerRequest(r WorkerRequest) error {

	creatworkerreqq.Do(func() {

		wexName := os.Getenv("WORKER_REQUEST_EXCHANGE")
		if wexName == "" {
			wexName = "worker_request_topic"
		}

		err := a.Ch.ExchangeDeclare(
			wexName, // name
			"topic", // type
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			log.Panic("unable to declare exchange for worker request topic")
		}
		a.WorkerRequestEx = wexName
	})

	req, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return a.sendRequest(req)
}

func (a *RabbitAMQPClient) sendRequest(p []byte) error {

	if a.WorkerRequestEx == "" {
		return errors.New("worker request exchange not found")
	}

	err := a.Ch.PublishWithContext(context.Background(),
		a.WorkerRequestEx, // exchange
		"",                // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})
	if err != nil {
		return err
	}
	return nil
}
