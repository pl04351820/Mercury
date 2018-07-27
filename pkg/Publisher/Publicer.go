package Publisher

import (
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/nats-io/go-nats"
)

type Publisher struct {
	Address string
	Svc     *nats.EncodedConn
}

type MsgType struct {
	InputEvents []byte
	Job         Type.Job
}

func NewPublisher(address string) Publisher {
	var client Publisher
	client.Address = address
	nc, _ := nats.Connect(address)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	client.Svc = c
	return client
}

func (p *Publisher) Public(job Type.Job, inputEvent []byte, subject string) {
	var me MsgType
	me.InputEvents = inputEvent
	me.Job = job

	res, _ := json.Marshal(me)
	p.Svc.Publish(subject, res)
}

func (p *Publisher) ParamCheck() {
	// TODO Implement param check.
	// Actually we can use open source tool from AWS

}
