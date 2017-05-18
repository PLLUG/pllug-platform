package main

import (
	"io"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

const (
	AMQTP_HOST = "amqp://guest:guest@192.168.99.100:5672/"
)

func Handler(w http.ResponseWriter, r *http.Request) {
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
	body := "Hello world"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal("Error publish message")
	}
	io.WriteString(w, "hello world\n")
}

func main() {
	http.HandleFunc("/", Handler)
	log.Println("App listen on http://192.168.99.100:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}