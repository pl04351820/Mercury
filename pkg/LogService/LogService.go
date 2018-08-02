package LogService

import (
	"Mercury/pkg/Type"
	"context"
	"github.com/olivere/elastic"
	"log"
	"Mercury/pkg/Conf"
)

type LogClient struct {
	Address string
	Ctx     context.Context
	Svc     *elastic.Client
}

type TaskLog struct {
	State string
	Next string
	Log string
}

type JobLog struct {
	JobName string
	Entry string
	TotalCount int
	Tasks map[string]TaskLog
}

func NewLogClient(Address string) LogClient {
	confObject := Conf.GetConf("conf.yaml")
	es, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(confObject.ElasticSearch))
	if err != nil {
		panic(err)
	}
	client := LogClient{Address:confObject.ElasticSearch, Svc:es, Ctx:context.Background()}
	return client
}

// Replace this implementation
func (l *LogClient) InitJobLog(job Type.Job) {
	totalCount := len(job.States)
	TasksState := make(map[string]TaskLog)
	for taskName, taskInfo := range job.States{
		newTask := TaskLog{State: "false", Next: taskInfo.Next, Log: ""}
		TasksState[taskName] = newTask
		}

	NewJobLog := JobLog{JobName:job.JobName, Entry:job.StartAt, TotalCount:totalCount, Tasks:TasksState}
	_, err := l.Svc.Index().Index("stepFunction").Type("log").BodyJson(NewJobLog).Do(l.Ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Init Job log state successfully")
}

//
func (l *LogClient) UpdateJobLog(taskName string) {

}
