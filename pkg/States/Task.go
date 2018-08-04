package States

import (
	"Mercury/pkg/Path"
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

func (t *Task) run() string {
	pathService := Path.NewJsonPathService(t.Common.Events)
	lambdaEvents := pathService.InputPathHandler(t.InputPath)

	lambdaResult, logResult, err := awsClient.Invoke(t.Resource, lambdaEvents)
	if err != nil {
		log.Printf("Error happen when invoke AWS function %v \n", err.Error())
	}

	pathService.ResultPathHandler(t.ResultPath, lambdaResult)

	t.Common.Events = pathService.OutputPathHandler(t.OutputPath)
	return logResult
}
