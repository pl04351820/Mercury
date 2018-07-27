package Subscriber

import (
	"Mercury/pkg/Type"
	"Mercury/plugin/AWS"
	"fmt"
	"github.com/nats-io/go-nats"
	//"Mercury/pkg/Log"
	"sync"
	"log"
	"time"
	"Mercury/pkg/Publisher"
	"encoding/json"
	"encoding/base64"
	"Mercury/pkg/LogService"
)

type Subscriber struct {
	NatAddress string
	Svc        *nats.Conn
}

var awsClient AWS.LambdaClient
var esClient LogService.LogClient

func init() {
	awsClient = AWS.NewLambdaClient()
	esClient = LogService.NewLogClient("http://0.0.0.0:9200/")
}

func NewSubscriber(address string) Subscriber {
	var client Subscriber
	client.NatAddress = address
	nc, err := nats.Connect(address)
	if err != nil{
		log.Fatalf("Error when connect to NATS %v ", err.Error())
	}
	//c, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	//if err != nil{
	//	log.Fatalf("Error when connect to NATS ENCODE %v", err.Error())
	//}
	client.Svc = nc
	return client
}

func (s *Subscriber) Subscribe(subject string, wg sync.WaitGroup) {
	// Infinite loop to handle goroutine]
	sub, err := s.Svc.SubscribeSync(subject)
	if err != nil {
		fmt.Println("Error when connect to nats ", err.Error())
	}

	for {
		msg, err := sub.NextMsg(time.Duration(6) * time.Second)
		if msg != nil{
			if err != nil{
				fmt.Printf("Error when stream %v", err.Error())
			}
			var res Publisher.MsgType
			err = json.Unmarshal(msg.Data, &res)
			if err != nil{
				fmt.Printf("Error when decode from json", err.Error())
			}
			ExecuteJob(&res.Job, res.InputEvents)
			time.Sleep(time.Duration(1) * time.Second)
		}

		// Use ASync Version
		//s.Svc.Subscribe(subject, func(msg *nats.Msg) {
		//	fmt.Println("GET NEW REQUEST #######################")
		//	var res Publisher.MsgType
		//	json.Unmarshal(msg.Data, &res)
		//	ExecuteJob(&res.Job, res.InputEvents)
		//	fmt.Println("GET NEW REQUEST END #######################")
		//})
	}
	wg.Done()
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

type ESType struct{
	TaskName string
	LogInfo string
}

func ExecuteTask(task Type.Task, input []byte) ([]byte, string, bool) {
	res, logResult, err := awsClient.Invoke(task.Resource, input)

	if err != nil{
		fmt.Println("Panic, error data %+v \n", err.Error())
	}
	// TODO: ADD LOG System into ElasticSearch


	decodeBytes, err := base64.StdEncoding.DecodeString(logResult)
	//fmt.Printf("The logResult is %v \n", string(decodeBytes))

	newLog := Type.ESType{TaskName:"new_step_function", LogInfo:string(decodeBytes)}
	esClient.InsertES("StepFunctionLog", newLog)

	if task.End == true {
		return res, "", true
	} else {
		return res, task.Next, false
	}
}
