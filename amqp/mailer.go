package amqp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MailRequest struct {
	From         string         `json:"from"`
	To           string         `json:"to"`
	Subject      string         `json:"subject"`
	DataMap      map[string]any `json:"data_map"`
	TemplateName string         `json:"template_name"`
	AckRequired  bool           `json:"ack_required"`
	Priority     string         `json:"priority"`
}

type MailResponse struct {
	Status        bool   `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlation_id"`
}

var createmailq sync.Once

func (r *RabbitAMQPClient) SendMail(msg []byte, domain string, cb func(MailResponse)) (string, error) {

	createmailq.Do(func() {

		q, err := r.Ch.QueueDeclare(
			"",    // name
			false, // durable
			true,  // delete when unused
			true,  // exclusive
			false, // noWait
			nil,   // arguments
		)
		if err != nil {
			log.Println("Unable to create mail queue ", err)
		}
		r.MailReqQName = domain
		r.MailResQName = q.Name
	})

	if r.MailReqQName == "" {
		return "", errors.New("unable to create mail queue")
	}

	msgs, err := r.Ch.Consume(
		r.MailResQName, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	if err != nil {
		log.Println("Failed to register a consumer ", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	corrId := randomString(32)

	err = r.Ch.PublishWithContext(ctx,
		"",             // exchange
		r.MailReqQName, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: corrId,
			ReplyTo:       r.MailResQName,
			Body:          msg,
		})
	if err != nil {
		log.Println("Failed to publish a message")
		return "", err
	}

	go func() {
		for d := range msgs {
			if corrId == d.CorrelationId {
				if cb != nil {
					var res MailResponse
					if err := json.Unmarshal(d.Body, &res); err != nil {
						log.Println("unble to invoke mail response callback function")
					}
					cb(res)
				}
				break
			}
		}
	}()

	fmt.Println("Mail request sent on queue", r.MailReqQName)

	return corrId, nil

}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
