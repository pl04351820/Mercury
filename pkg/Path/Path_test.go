package Path

import (
	"testing"
	"encoding/json"
	"log"
	"bytes"
	"k8s.io/client-go/util/jsonpath"
	//_"github.com/mdaverde/jsonpath"
)

func TestNewJsonPathService(t *testing.T) {
	data := []byte(`{"expensive":"Result"}`)
	var json_data interface{}
	json.Unmarshal([]byte(data), &json_data)
	// testJsonService := NewJsonPathService(json_data)

	// Read from jsonPath
	// res := testJsonService.GetDataFromJsonPath("$.expensive")
	//log.Println(res)

	// Write to jsonPath


	// Test using Kubernetes jsonPath

	j := jsonpath.New("TheJsonTest")
	err := j.Parse("{$.expensive}")
	if err != nil{
		log.Fatal("Error Happens")
	}
	//fullResults, err := j.FindResults(json_data)
	//log.Println(fullResults)
	//log.Println(fullResults[0][0])


	buf := new(bytes.Buffer)
	err = j.Execute(buf, json_data)

	//fullResults, err = j.FindResults(json_data)
	//log.Println(fullResults[0][0])

	out := buf.String()
	log.Println(out)
}