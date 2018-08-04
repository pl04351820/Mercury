package Publisher

import (
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"log"
)

type Publisher struct {
	Address string
	Svc     *nats.Conn
}

type MsgType struct {
	InputEvents []byte
	Job         Type.Job
}

func NewPublisher(address string) Publisher {
	nc, err := nats.Connect(address)
	if err != nil {
		log.Fatalf("The NATS connection is failed %v", err.Error())
	}
	client := Publisher{Address: address, Svc: nc}
	return client
}

func (p *Publisher) Public(job Type.Job, inputEvent []byte, subject string) {
	me := MsgType{InputEvents: inputEvent, Job: job}
	res, err := json.Marshal(me)
	if err != nil {
		log.Fatalf("Marshall error when public %v", err.Error())
	}

	err = p.Svc.Publish(subject, res)
	if err != nil {
		log.Fatalf("Error when publishing message %v", err.Error())
	}
	log.Println("Public Successfully")
}

func (p *Publisher) ParamCheck() {
	// TODO Implement param check.
}
