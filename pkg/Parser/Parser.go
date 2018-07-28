package Parser

import (
	"Mercury/pkg/Type"
	"encoding/json"
	"io/ioutil"
	"log"
)

func ParserJob(filename string) Type.Job {
	var job Type.Job
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("JsonFile.Get err #%v ", err)
	}
	err = json.Unmarshal(file, &job)
	if err != nil {
		log.Fatal("Unmarshal: %v", err)
	}
	return job
}

func ParseEvents(filename string) []byte {
	// For now, the events can only use string as its key.
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Events file get err #%v", err)
	}
	return file
}
