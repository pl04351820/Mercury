package Subscriber

import (
	"Mercury/pkg/Publisher"
	"Mercury/pkg/States"
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"log"
	"sync"
	"time"
)

type Subscriber struct {
	NatAddress string
	Svc        *nats.Conn
}

func NewSubscriber(address string) Subscriber {
	nc, err := nats.Connect(address)
	if err != nil {
		log.Fatalf("Error when connect to NATS %v ", err.Error())
	}
	client := Subscriber{Svc:nc, NatAddress:address}
	return client
}

func (s *Subscriber) Subscribe(subject string, wg sync.WaitGroup) {
	sub, err := s.Svc.SubscribeSync(subject)
	if err != nil {
		log.Fatalf("Error when connect to nats %v ", err.Error())
	}
	// TODO: Init Log Status.

	for {
		msg, err := sub.NextMsg(time.Duration(6) * time.Second)
		if msg != nil {
			if err != nil {
				log.Printf("Error when stream %v", err.Error())
			}
			var res Publisher.MsgType
			err = json.Unmarshal(msg.Data, &res)
			if err != nil {
				log.Printf("Error when decode from json %v", err.Error())
			}
			ExecuteJob(&res.Job, res.InputEvents)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
	wg.Done()
}

func ExecuteJob(job *Type.Job, input []byte) {
	curState := job.StartAt
	stateQueue := make([]string, 30)
	stateQueue = append(stateQueue, curState)
	for len(stateQueue) > 0 {
		curState = stateQueue[0]
		stateQueue = stateQueue[1:]
		nextState, output := States.TaskState(job.States[curState], input)
		input = output
		// TODO Update ES

		if nextState != "" {
			stateQueue = append(stateQueue, nextState)
		}
	}
}