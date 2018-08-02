package States

import (
	"Mercury/pkg/LogService"
	"Mercury/pkg/Type"
	"Mercury/plugin/AWS"
	"encoding/base64"
	"log"
)


type Task struct {
	Common CommonField
	Resource string `json:"Resource"`
	Next string `json:"Next"`
	ResultPath string `json:"ResultPath"`
	End bool `json:"End"`
}






// Legacy
var awsClient AWS.LambdaClient
var esClient LogService.LogClient

func init() {
	awsClient = AWS.NewLambdaClient()
	esClient = LogService.NewLogClient("http://0.0.0.0:9200/")
}

func TaskState(task Type.Task, events []byte) (string, []byte) {
	res, logResult, err := awsClient.Invoke(task.Resource, events)
	if err != nil {
		log.Printf("Error happen when invoke AWS function %v \n", err.Error())
	}

	decodeBytes, err := base64.StdEncoding.DecodeString(logResult)
	newLog := Type.ESType{TaskName: "new_step_function", LogInfo: string(decodeBytes)}
	esClient.InsertES(newLog)
	return task.Next, res
}
