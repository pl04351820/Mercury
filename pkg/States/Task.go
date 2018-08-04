package States

import (
	"Mercury/plugin/AWS"
	"log"
)

type Task struct {
	Common     CommonField
	Resource   string `json:"Resource"`
	Next       string `json:"Next"`
	End        bool   `json:"End"`
	ResultPath string `json:"ResultPath"`
	InputPath  string `json:"InputPath"`
	OutputPath string `json:"OutputPath"`
}

var awsClient AWS.LambdaClient

func init() {
	awsClient = AWS.NewLambdaClient()
	//esClient = LogService.NewLogClient("http://0.0.0.0:9200/")
}

// Log should be implemented outside the attribute.
// Path should be handled in the method instead of out including the events.


func (t *Task) run() string {
	

	_, logResult, err := awsClient.Invoke(t.Resource, t.Common.Events)
	if err != nil {
		log.Printf("Error happen when invoke AWS function %v \n", err.Error())
	}
	return logResult
}
