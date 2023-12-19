package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

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

type WorkerResponse struct {
	RequestId         string   `json:"request_id"`
	RequestType       string   `json:"request_type"`
	IsProcessed       bool     `json:"is_processed"`
	IsError           bool     `json:"is_error"`
	OriginalSize      int64    `json:"original_size_in_bytes,omitempty"`
	OptimizedSize     int64    `json:"optimized_size_in_bytes,omitempty"`
	CompressionRatio  float64  `json:"compression_ratio,omitempty"`
	ProcessingTime    float64  `json:"processing_time_in_ms,omitempty"`
	OutputFileName    string   `json:"output_file_name,omitempty"`
	OutputZipFileName string   `json:"output_zip_file_name,omitempty"`
	OutputFileList    []string `json:"output_file_list,omitempty"`
	MergedFileSize    int64    `json:"merged_file_size,omitempty"`
	Message           string   `json:"message,omitempty"`
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
			fmt.Println("Unable to create worker queue ", err)
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
		fmt.Println("Failed to register a consumer ", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Minute)
	defer cancel()

	fmt.Println("Publishing worker request" + " " + string(msg))

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
		fmt.Println("Failed to publish a message")
		return "", err
	}

	go func() {
		for d := range msgs {
			fmt.Println("received response")
			if requestId == d.CorrelationId {
				if cb != nil {
					var res WorkerResponse
					if err := json.Unmarshal(d.Body, &res); err != nil {
						fmt.Println("unble to invoke worker response callback function")
					}
					cb(res)
				}
				break
			}
		}
	}()

	return requestId, nil

}
