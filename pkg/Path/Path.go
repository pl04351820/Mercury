package Path

/*
Support for Path method with aws protocol. The Path is represented by []byte.

InputPath
ResultPath
OutputPath

In this stage, since the golang jsonPath is short of retrieving option.
We use another name style, adn we will change it later.
*/

import (
	_ "encoding/json"
	"github.com/mdaverde/jsonpath"
	"log"
)

type JsonPathService struct {
	JsonData interface{}
	RawData []byte
}

func NewJsonPathService(JsonData interface{}, RawData []byte) JsonPathService {
	
	return JsonPathService{JsonData: JsonData, RawData: RawData}
}

func (j *JsonPathService) GetDataFromJsonPath(path string) interface{} {
	value, err := jsonpath.Get(j.JsonData, path)
	if err != nil {
		log.Fatalf("Read error from jsonpath %s", value)
	}
	return value
}

func (j *JsonPathService) WriteDataToJsonPath(path string, content interface{}) {
	err := jsonpath.Set(j.JsonData, path, content)
	if err != nil {
		log.Fatal(err)
	}
}

func (j *JsonPathService) GetDataFromInputPath(inputPath string) {

}
