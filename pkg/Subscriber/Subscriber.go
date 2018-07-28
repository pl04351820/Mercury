package Subscriber

import (
	"Mercury/pkg/Publisher"
	"Mercury/pkg/States"
	"Mercury/pkg/Type"
	"encoding/json"
	"fmt"
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
	var client Subscriber
	client.NatAddress = address
	nc, err := nats.Connect(address)
	if err != nil {
		log.Fatalf("Error when connect to NATS %v ", err.Error())
	}
	client.Svc = nc
	return client
}

func (s *Subscriber) Subscribe(subject string, wg sync.WaitGroup) {
	sub, err := s.Svc.SubscribeSync(subject)
	if err != nil {
		fmt.Println("Error when connect to nats ", err.Error())
	}

	for {
		msg, err := sub.NextMsg(time.Duration(6) * time.Second)
		if msg != nil {
			if err != nil {
				fmt.Printf("Error when stream %v", err.Error())
			}
			var res Publisher.MsgType
			err = json.Unmarshal(msg.Data, &res)
			if err != nil {
				fmt.Printf("Error when decode from json %v", err.Error())
			}
			ExecuteJob(&res.Job, res.InputEvents)
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
	wg.Done()
}

func ExecuteJob(job *Type.Job, input []byte) {
	// Add to es for first state

	curState := job.StartAt
	stateQueue := make([]string, 30)
	stateQueue = append(stateQueue, curState)
	for len(stateQueue) > 0 {
		curState = stateQueue[0]
		stateQueue = stateQueue[1:]
		nextState, output := States.TaskState(job.States[curState], input)
		input = output

		// Add to es for every loop

		if nextState != "" {
			stateQueue = append(stateQueue, nextState)
		}
	}
}
