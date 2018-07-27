package main

import (
	"Mercury/pkg/Subscriber"
	"Mercury/pkg/Conf"
	"Mercury/pkg/Publisher"
	"Mercury/pkg/Parser"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	// Read configuration from file.
	confTarget := Conf.GetConf()

	// Read job and input events.
	job := Parser.ParserJob("demo.json")
	inputEvents := Parser.ParseEvents("inputEvents.json")

	// Create Subscriber
	subscriber := Subscriber.NewSubscriber(confTarget.Nats)
	subscriber.Subscribe("foo", wg)

	publisher := Publisher.NewPublisher(confTarget.Nats)
	publisher.Public(job, inputEvents, "foo")

	wg.Wait()
}