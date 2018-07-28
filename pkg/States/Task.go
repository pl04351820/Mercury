package States

import (
	"Mercury/pkg/LogService"
	"Mercury/pkg/Type"
	"Mercury/plugin/AWS"
	"encoding/base64"
	"fmt"
)

var awsClient AWS.LambdaClient
var esClient LogService.LogClient

func init() {
	awsClient = AWS.NewLambdaClient()
	esClient = LogService.NewLogClient("http://0.0.0.0:9200/")
}

func TaskState(task Type.Task, events []byte) (string, []byte) {
	res, logResult, err := awsClient.Invoke(task.Resource, events)
	if err != nil {
		fmt.Printf("Error happen when invoke AWS function %v \n", err.Error())
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(logResult)
	newLog := Type.ESType{TaskName: "new_step_function", LogInfo: string(decodeBytes)}
	esClient.InsertES("StepFunctionLog", newLog)
	return task.Next, res
}