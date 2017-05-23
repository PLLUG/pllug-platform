package main

import (
	"log"
	"os"
	"github.com/streadway/amqp"
)

func main() {
	log.Println("connect to host: %s", os.Getenv("AMQP_HOST"))
	conn, err := amqp.Dial(os.Getenv("AMQP_HOST"))
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
	messages, err := ch.Consume(
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
		for message := range messages {
			log.Printf("Received a message: %s", message.Body)
		}
	}()
	<-forever
}

func failOnError(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}