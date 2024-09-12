package main

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	// connect to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Printf("failure to connect: %s", err)
	}
	defer conn.Close()

	// connect to rabbitmq's channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Printf("failure to connect channel: %s", err)
	}
	defer ch.Close()

	// declare queue
	q, err := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Printf("failure to declare queue: %s", err)
	}

	type User struct {
		Email    string
		Password string
	}

	user := User{
		Email:    "abimo@gmail.com",
		Password: "123123123",
	}

	dataJson, _ := json.Marshal(user)

	// publish queue
	if err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(dataJson),
		},
	); err != nil {
		fmt.Printf("fail to publish queue: %s", err)
	}
	fmt.Println("Successfully to publish queue")
}
