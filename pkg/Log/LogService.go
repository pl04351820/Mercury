package LogService

import "github.com/olivere/elastic"

type LogClient struct{
	Address string

}

func (l *LogClient)