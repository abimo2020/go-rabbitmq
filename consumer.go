package main

import (
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

	// consume messages
	msg, err := ch.Consume(
		"test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msg {
			fmt.Printf("Receive message: %s", d.Body)
		}
	}()
	fmt.Println("[*] waiting for message..")
	<-forever
}
