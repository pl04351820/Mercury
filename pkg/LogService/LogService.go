package LogService

import (
	"Mercury/pkg/Type"
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type LogClient struct {
	Address string
	Ctx     context.Context
	Svc     *elastic.Client
}

func NewLogClient(Address string) LogClient {
	var client LogClient
	client.Ctx = context.Background()
	// Connect to ElasticSearch, turn off sniff when you use container.
	es, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200/"))
	fmt.Println("Before Panic")
	if err != nil {
		panic(err)
	}
	fmt.Println("After Panic")
	client.Svc = es
	return client
}

func (l *LogClient) InsertES(content Type.ESType) {
	_, err := l.Svc.Index().Index("stepFunction").Type("log").BodyJson(content).Do(l.Ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert INTO ES successfully!")
}

func (l *LogClient) InitJobLog(job Type.Job){
	esJobState := Type.JobState{}
	stateMap := make(map[string]bool)
	for taskName, _ := range(job.States){
		stateMap[taskName] = false
	}
	esJobState.StatueInfo = stateMap
	esJobState.JobName = "NewJob"
	_, err := l.Svc.Index().Index("stepFunction").Type("log").BodyJson(esJobState).Do(l.Ctx)
	if err != nil{
		panic(err)
	}
	fmt.Println("Init Job State successfully!")
}

func (l *LogClient) UpdateJobLog(taskName string){
	// Read from Es and write to it.
}