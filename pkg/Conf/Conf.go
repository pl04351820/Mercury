package Conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	ElasticSearch string `yaml:"Elasticsearch"`
	Nats          string `yaml:"Nats"`
}

func GetConf(filename string) Conf {
	c := Conf{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("YamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}

	return c
}
