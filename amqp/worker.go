package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerRequest struct {
	RequestId   string   `json:"request_id"`
	RequestType string   `json:"request_type"`
	PageNumbers []int    `json:"page_numbers,omitempty"`
	Ranges      []string `json:"ranges,omitempty"`
}

type WorkerResponse struct {
	Status      bool   `json:"status"`
	RequestId   string `json:"request_id"`
	RequestType string `json:"request_type"`
	Message     string `json:"message"`
}

var creatworkerq sync.Once

func (r *RabbitAMQPClient) SendWorkRequest(msg []byte, requestId string, cb func(WorkerResponse)) (string, error) {

	creatworkerq.Do(func() {

		rqname := os.Getenv("WORKER_QUEUE_NAME")
		if rqname == "" {
			rqname = "worker_queue"
		}

		q, err := r.Ch.QueueDeclare(
			"",    // name
			false, // durable
			true,  // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			log.Println("Unable to create worker queue ", err)
		}
		r.WorkerReqQName = rqname
		r.WorkerResQName = q.Name
	})

	if r.WorkerReqQName == "" {
		return "", errors.New("unable to create worker queue")
	}

	msgs, err := r.Ch.Consume(
		r.WorkerResQName, // queue
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Println("Failed to register a consumer ", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = r.Ch.PublishWithContext(ctx,
		"",               // exchange
		r.WorkerReqQName, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: requestId,
			ReplyTo:       r.WorkerResQName,
			Body:          msg,
		})
	if err != nil {
		log.Println("Failed to publish a message")
		return "", err
	}

	go func() {
		for d := range msgs {
			if requestId == d.CorrelationId {
				if cb != nil {
					var res WorkerResponse
					if err := json.Unmarshal(d.Body, &res); err != nil {
						log.Println("unble to invoke worker response callback function")
					}
					cb(res)
				}
				break
			}
		}
	}()

	return requestId, nil

}
