package Publisher

// Import packages
import (
	"Mercury/pkg/Type"
	"github.com/nats-io/go-nats"
	"encoding/json"
	"log"
)

type Publisher struct{
	NatsAddress string
}

func (p *Publisher)Public(job Type.Job, conn nats.Conn, subject string){
	// TODO: Parameter Check

	res, _ := json.Marshal(job)
	conn.Publish(subject, res)
	log.Println("Published Message on subject" + subject)

}