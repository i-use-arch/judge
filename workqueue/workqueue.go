package workqueue

import (
	"context"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

// QFunc is a function that takes a queue message.
type QFunc func(uint64) error

// MakeQueue creates a connection to the work queue.
func MakeQueue(ctx context.Context, rabbitStr string, f QFunc) {
	conn, err := amqp.Dial(rabbitStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	msgs, err := ch.Consume(
		"submission_queue", // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)

	go func() {
		for d := range msgs {
			log.Printf("new message: %s\n", d.Body)
			id, err := strconv.ParseUint(string(d.Body), 10, 64)
			if err != nil {
				log.Printf("%v", err)
				continue
			}
			err = f(id)
			if err != nil {
				log.Printf("%v", err)
			}
		}
	}()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
