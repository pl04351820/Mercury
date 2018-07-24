package LogService

import (
	"context"
	"github.com/olivere/elastic"
)

type LogClient struct{
	Address string
	Ctx context.Context
	Svc *elastic.Client
}

func NewLogClient(Address string) LogClient{
	var client LogClient
	client.Ctx = context.Background()
	// Connect to
	es, err := elastic.NewClient(elastic.SetURL(Address))
	if err != nil{
		panic(err)
	}
	client.Svc = es
	return client
}

func (l *LogClient) InsertES(){
	//
}