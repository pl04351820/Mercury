package Subscriber

import (
	"github.com/nats-io/go-nats"
	"fmt"
	"Mercury/pkg/Type"
	"Mercury/plugin/AWS"
)

type Subscriber struct{
	NatsAddress string
}

var awsClient AWS.LambdaClient

func init(){
	awsClient = AWS.NewLambdaClient()
}

func (s *Subscriber) Subscribe(conn nats.Conn, subject string){
	// This is the element unit for goroutine.

	// Infinite loop to handle goroutine
	for {
		conn.Subscribe(subject, func(msg *nats.Msg){
			// UnMarshall


			conn.Publish(msg.Reply, []byte("I can help"))
		})
	}
}

func ExecuteJob(job *Type.Job, input []byte){
	curState := job.StartAt
	endFlag := false
	var output []byte
	var nextState string
	for endFlag == false{
		output, nextState, endFlag = ExecuteTask(job.States[curState], input)
		input = output
		curState = nextState
	}
}

func ExecuteTask(task Type.Task, input []byte) ([]byte, string, bool){
	// TODO: ADD LOG System into ElasticSearch
	res, err := awsClient.Invoke(task.Resource, input)
	if err != nil{
		fmt.Println(err.Error())
	}

	if task.End == true{
		return res, "", true
	}else{
		return res, task.Next, false
	}
}