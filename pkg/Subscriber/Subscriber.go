package Subscriber

import (
	"Mercury/pkg/LogService"
	"Mercury/pkg/Parser"
	"Mercury/pkg/Publisher"
	"Mercury/pkg/States"
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/nats-io/go-nats"
	"log"
	"reflect"
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
	client := Subscriber{Svc: nc, NatAddress: address}
	return client
}

/*
Step1: Init Log status and init job configuration.
Step2: Run FSM loop
*/

func (s *Subscriber) Subscribe(subject string, wg sync.WaitGroup) {
	sub, err := s.Svc.SubscribeSync(subject)
	defer wg.Done()
	if err != nil {
		log.Fatalf("Error when connect to nats %v ", err.Error())
	}
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
}

func (s *Subscriber) executeJob(jobMsg []byte) {
	// Init log service state
	logSvc := LogService.NewLogClient()
	logSvc.InitJobLog(jobMsg)
	parser := Parser.NewParser(jobMsg)
	curStateName := parser.PathSvc.GetDataFromJsonPath(".StartAt").(string)
	curState := parser.ParseTask(curStateName)

	for {
		var nextState string
		var logText string
		switch curState.(type) {
		case States.Task:
			taskState := curState.(States.Task)
			logText = taskState.Run()
			nextState = taskState.Next
		case States.Pass:
			passState := curState.(States.Pass)
			logText = passState.Run()
			nextState = passState.Next
		case States.Choice:
			choiceState := curState.(States.Choice)
			logText = choiceState.Run()
			nextState = choiceState.Next
		case States.Wait:
			waitState := curState.(States.Wait)
			logText = waitState.Run()
			nextState = waitState.Next
		}
		logSvc.UpdateJobLog(curStateName, logText)
		if nextState == "" {
			break
		}
		curStateName = nextState
		curState = parser.ParseTask(curStateName)
	}
}

//func ExecuteJob(job *Type.Job, input []byte) {
//	curState := job.StartAt
//	stateQueue := make([]string, 30)
//	stateQueue = append(stateQueue, curState)
//	for len(stateQueue) > 0 {
//		curState = stateQueue[0]
//		stateQueue = stateQueue[1:]
//		nextState, output := States.TaskState(job.States[curState], input)
//		input = output
//		// TODO Update ES
//
//		if nextState != "" {
//			stateQueue = append(stateQueue, nextState)
//		}
//	}
//}
