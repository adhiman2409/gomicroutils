package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerRequest struct {
	RequestId   string   `json:"request_id"`
	RequestType string   `json:"request_type"`
	Domain      string   `json:"domain,omitempty"`
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

		q, err := a.Ch.QueueDeclare(
			wexName, // name
			true,    // durable
			false,   // auto-deleted
			false,   // internal
			false,   // no-wait
			nil,     // arguments
		)
		if err != nil {
			log.Panic("unable to declare exchange for worker request topic")
		}
		a.WorkerRequestQ = q.Name
	})

	req, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return a.sendRequest(req)
}

func (a *RabbitAMQPClient) sendRequest(p []byte) error {

	if a.WorkerRequestQ == "" {
		fmt.Println("worker request queue not found")
		return errors.New("worker request queue not found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := a.Ch.PublishWithContext(ctx,
		"",               // exchange
		a.WorkerRequestQ, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         p,
		})
	if err != nil {
		fmt.Println("worker request publish error")
		return err
	}
	fmt.Println("worker request send successfully")
	return nil
}
