package Path

/*
Support for Path method with aws protocol. The Path is represented by []byte.

The jsonPath rule:

InputPath: If InputPath == "":
				Pass the whole events to lambda function
           else:
				Pass the jsonPath

ResultPath: If ResultPath == "":
				Do not change anything
			else:
				Change specific field.

OutputPath: If OutputPath == "":
				Pass all output to Next Field
			else:
				Pass the jsonPath

In this stage, since the golang jsonPath is short of retrieving option.
We use another name style, adn we will change it later.
*/

import (
	"encoding/json"
	"github.com/mdaverde/jsonpath"
	"log"
)

type JsonPathService struct {
	JsonData interface{}
}

func NewJsonPathService(RawData []byte) JsonPathService {
	var json_data interface{}
	json.Unmarshal([]byte(RawData), &json_data)
	return JsonPathService{JsonData: json_data}
}

func (j *JsonPathService) GetDataFromJsonPath(path string) interface{} {
	value, err := jsonpath.Get(j.JsonData, path)
	if err != nil {
		log.Fatalf("Read error from jsonpath %v", err)
	}
	return value
}

func (j *JsonPathService) WriteDataToJsonPath(path string, content interface{}) {
	err := jsonpath.Set(j.JsonData, path, content)
	if err != nil {
		log.Fatal(err)
	}
}

func (j *JsonPathService) InputPathHandler(inputPath string) []byte {
	var source interface{}
	if inputPath == "" {
		source = j.JsonData
	} else {
		source = j.GetDataFromJsonPath(inputPath)
	}

	res, err := json.Marshal(source)
	if err != nil {
		log.Fatalf("Error for marshall json %s", err)
	}
	return res
}

func (j *JsonPathService) ResultPathHandler(resultPath string, content interface{}) {
	if resultPath == "" {
		return
	} else {
		j.WriteDataToJsonPath(resultPath, content)
	}
}

func (j *JsonPathService) OutputPathHandler(outputPath string) []byte {
	var source interface{}
	if outputPath == "" {
		source = j.JsonData
	} else {
		source = j.GetDataFromJsonPath(outputPath)
	}

	res, err := json.Marshal(source)
	if err != nil {
		log.Fatalf("Error for marshall json %s", err)
	}
	return res
}
