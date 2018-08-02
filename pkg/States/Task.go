package States

import (
	"Mercury/plugin/AWS"
	"log"
)


type Task struct {
	Common CommonField
	Resource string `json:"Resource"`
	Next string `json:"Next"`
	ResultPath string `json:"ResultPath"`
	End bool `json:"End"`
}

var awsClient AWS.LambdaClient

func init() {
	awsClient = AWS.NewLambdaClient()
	//esClient = LogService.NewLogClient("http://0.0.0.0:9200/")
}

// Log should be implemented outside the attribute.
// Path should be handled in the method instead of out including the events.
func (t *Task) run()(string){
	_, logResult, err := awsClient.Invoke(t.Resource, t.Common.Events.([]byte))
	if err != nil {
		log.Printf("Error happen when invoke AWS function %v \n", err.Error())
	}
	return logResult
}
