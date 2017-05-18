package main

import (
	"log"

	"github.com/streadway/amqp"
)

const (
	AMQTP_HOST = "amqp://guest:guest@192.168.99.100:5672/"
)

func main() {
	conn, err := amqp.Dial(AMQTP_HOST)
	if err != nil {
		log.Fatal("%s", err)
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Error connection to channel")
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"test_q",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error create queue")
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error consume")
	}

forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-forever;
}