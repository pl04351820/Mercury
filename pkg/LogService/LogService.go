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

func NewLogClient(Address string) LogClient {
	confObject := Conf.GetConf("conf.yaml")
	es, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(confObject.ElasticSearch))
	if err != nil {
		panic(err)
	}
	client := LogClient{Address:confObject.ElasticSearch, Svc:es, Ctx:context.Background()}
	return client
}

func (l *LogClient) InsertES(content Type.ESType) {
	_, err := l.Svc.Index().Index("stepFunction").Type("log").BodyJson(content).Do(l.Ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Insert INTO ES successfully")
}

func (l *LogClient) InitJobLog(job Type.Job) {
	stateMap := make(map[string]bool)
	for taskName, _ := range job.States {
		stateMap[taskName] = false
	}
	// TODO: Replace JobName
	esJobState := Type.JobState{StatueInfo:stateMap, JobName:"NewJob"}
	_, err := l.Svc.Index().Index("stepFunction").Type("log").BodyJson(esJobState).Do(l.Ctx)
	if err != nil {
		panic(err)
	}
	log.Println("Init Job log state successfully")
}

func (l *LogClient) UpdateJobLog(taskName string) {
}
