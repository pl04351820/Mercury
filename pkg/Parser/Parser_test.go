package Parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParser(t *testing.T) {
	data, err := ioutil.ReadFile("demo.json")
	if err != nil {
		log.Fatal(err)
	}
	parserSvc := NewParser(data)
	log.Printf("%+v", parserSvc.ParseTask("FirstState"))
	log.Printf("%+v", parserSvc.ParseTask("ChoiceState"))
	log.Printf("%+v", parserSvc.ParseTask("WarningToEmail"))
	log.Printf("%+v", parserSvc.ParseTask("DummyPass"))
}
