package main

import (
	"log"

	"github.com/streadway/amqp"
)

const (
	AMQP_HOST = "amqp://guest:guest@192.168.99.100:5672/"
)

func main() {
	conn, err := amqp.Dial(AMQP_HOST)
	failOnError(err, "Error connect to amqp server:")
	defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Error connection to channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"test_q",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Error create queue")
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Error consume")

forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever
}

func failOnError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}