package Receiver

// Import Go and NATS packages
import (
	"log"
	"github.com/nats-io/go-nats"
	"runtime"
)

func Receiver() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	log.Printf("Subscribing to subject 'foo'\n")
	natsConnection.Subscribe("foo", func(msg *nats.Msg) {
		// Handle the message
		log.Printf("Received message '%s\n", string(msg.Data)+"'")
	})

	// Keep the connection alive
	runtime.Goexit()
}