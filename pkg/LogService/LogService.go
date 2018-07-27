package LogService

import (
	"context"
	"github.com/olivere/elastic"
	"fmt"
	"Mercury/pkg/Type"
)

type LogClient struct{
	Address string
	Ctx context.Context
	Svc *elastic.Client
}

func NewLogClient(Address string) LogClient{
	var client LogClient
	client.Ctx = context.Background()
	// Connect to ElasticSearch, turn off sniff when you use container.
	es, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:9200/"))
	fmt.Println("Before Panic")
	if err != nil{
		panic(err)
	}
	fmt.Println("After Panic")
	client.Svc = es
	return client
}


func (l *LogClient) InsertES(indexName string, content Type.ESType){
	// Add Log to Elastic search

	_, err := l.Svc.Index().Index("stepfunction").Type("log").BodyJson(content).Do(l.Ctx)
	if err != nil{
		panic(err)
	}
	fmt.Println("Insert INTO ES successfully!")
}