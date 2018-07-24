package Publicer

// Import packages
import (
		"log"
		"github.com/nats-io/go-nats"
)

func Publicer() {
	// Connect to server; defer close
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// Publish message on subject
	subject := "foo"
	natsConnection.Publish(subject, []byte("Hello Message 2333"))
	log.Println("Published message on subject " + subject)
}