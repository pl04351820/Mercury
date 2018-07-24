package Subscriber

import (
	"Mercury/pkg/Publisher"
	"Mercury/pkg/Type"
	"Mercury/plugin/AWS"
	"encoding/json"
	"fmt"
	"github.com/nats-io/go-nats"
)

type Subscriber struct {
	NatAddress string
	Svc        *nats.EncodedConn
}

var awsClient AWS.LambdaClient

func init() {
	awsClient = AWS.NewLambdaClient()
}

func NewSubscriber(address string) Subscriber {
	var client Subscriber
	client.NatAddress = address
	nc, _ := nats.Connect(address)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	client.Svc = c
	return client
}

func (s *Subscriber) Subscribe(subject string) {
	// Infinite loop to handle goroutine
	for {
		s.Svc.Subscribe(subject, func(msg *nats.Msg) {
			// UnMarshall
			var res Publisher.MsgType
			json.Unmarshal(msg.Data, &res)
			ExecuteJob(&res.Job, res.InputEvents)
		})
	}
}

func ExecuteJob(job *Type.Job, input []byte) {
	curState := job.StartAt
	endFlag := false
	var output []byte
	var nextState string
	for endFlag == false {
		output, nextState, endFlag = ExecuteTask(job.States[curState], input)
		input = output
		curState = nextState
	}
}

func ExecuteTask(task Type.Task, input []byte) ([]byte, string, bool) {
	res, log, err := awsClient.Invoke(task.Resource, input)
	// TODO: ADD LOG System into ElasticSearch

	if err != nil {
		fmt.Println(err.Error())
	}

	if task.End == true {
		return res, "", true
	} else {
		return res, task.Next, false
	}
}
