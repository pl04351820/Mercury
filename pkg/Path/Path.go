package Path

/*
Support for Path method with aws protocol. The Path is represented by []byte.

InputPath
ResultPath
OutputPath
 */

 import (
 	_"github.com/oliveagle/jsonpath"
 	_"encoding/json"
	 "github.com/oliveagle/jsonpath"
	 "log"
 )

type JsonPathService struct {
	JsonData interface{}
}

func NewJsonPathService (RawJson interface{}) (JsonPathService){
	return JsonPathService{JsonData:RawJson}
}

func (j *JsonPathService) GetDataFromJsonPath(path string)(interface{}){
	res, err := jsonpath.JsonPathLookup(j.JsonData, path)
	if err != nil{
		log.Fatal("Error Happen when using json")
	}
	return res
}

func (j *JsonPathService) WriteDataToJsonPath(path string, content interface{}){

}



