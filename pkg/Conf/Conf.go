package Conf

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	ElasticSearch string `yaml:"elasticsearch"`
	Nats string `yaml:"nats"`
}

func GetConf() *Conf{
	var c *Conf
	file, err := ioutil.ReadFile("conf.yaml")
	if err != nil{
		log.Printf("YamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(file, c)
	if err != nil{
		log.Fatal("Unmarshal: %v", err)
	}
	return c
}