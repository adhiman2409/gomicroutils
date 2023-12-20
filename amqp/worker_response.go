package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

func (a *RabbitAMQPClient) SendWorkerResponse(r WorkerResponse) error {

	creatworkerq.Do(func() {

		wexName := os.Getenv("WORKER_RESPONSE_EXCHANGE")
		if wexName == "" {
			wexName = "worker_response_topic"
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
			fmt.Println("unable to declare exchange for worker response topic")
		}
		a.WorkerResponseEx = wexName
	})

	req, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error in marshaling message " + err.Error())
		return err
	}

	return a.sendResponse(req)
}

func (a *RabbitAMQPClient) sendResponse(p []byte) error {

	if a.WorkerResponseEx == "" {
		fmt.Println("worker response exchange not found")
		return errors.New("worker response exchange not found")
	}

	err := a.Ch.PublishWithContext(context.Background(),
		a.WorkerResponseEx, // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        p,
		})
	if err != nil {
		fmt.Println("worker response publish error " + err.Error())
		return err
	}
	fmt.Println("worker response sent successfully")
	return nil
}
