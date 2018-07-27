package Publisher

import (
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"fmt"
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
	var client Publisher
	client.Address = address
	nc, err := nats.Connect(address)
	if err != nil{
		fmt.Printf("The NATS connection is failed %v", err.Error())
	}
	//c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	//if err != nil{
	//	fmt.Printf("The NATS encode connection is failed %v", err.Error())
	//}
	client.Svc = nc
	return client
}

func (p *Publisher) Public(job Type.Job, inputEvent []byte, subject string) {
	var me MsgType
	me.InputEvents = inputEvent
	me.Job = job
	res, err := json.Marshal(me)

	if err != nil{
		fmt.Printf("Marshall error when public %v", err.Error())
	}

	err = p.Svc.Publish(subject, res)
	if err != nil{
		fmt.Printf("Error when publishing message %v", err.Error())
	}
	fmt.Println("Public Successfully")
}

func (p *Publisher) ParamCheck() {
	// TODO Implement param check.
	// Actually we can use open source tool from AWS

}
