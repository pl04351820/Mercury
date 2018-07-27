package main

import (
	"Mercury/pkg/Conf"
	"Mercury/pkg/Publisher"
	"Mercury/pkg/Parser"
	"sync"
	"Mercury/pkg/Subscriber"
	"fmt"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	// Read configuration from file.
	confTarget := Conf.GetConf("../conf.yaml")

	// Read job and input events.
	job := Parser.ParserJob("../demo.json")
	fmt.Printf("The input after decoding is %+v \n", job)
	inputEvents := Parser.ParseEvents("../inputEvents.json")
	fmt.Printf("The input after decoding is %+v \n", inputEvents)
	// Create Subscriber
	subscriber := Subscriber.NewSubscriber(confTarget.Nats)
	go subscriber.Subscribe("foo", wg)

	publisher := Publisher.NewPublisher(confTarget.Nats)
	publisher.Public(job, inputEvents, "foo")
	wg.Wait()
}