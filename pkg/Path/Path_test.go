package Path

import (
	"encoding/json"
	"log"
	"testing"
)

func TestNewJsonPathService(t *testing.T) {
	testData := []byte(`{"expensive":{"Result":"Art"}}`)
	var json_data interface{}
	json.Unmarshal([]byte(testData), &json_data)
	testJsonService := NewJsonPathService(json_data)

	// Read from jsonPath
	res := testJsonService.GetDataFromJsonPath("expensive.Result")
	log.Println(res)

	// Retrieve to jsonPath
	testJsonService.WriteDataToJsonPath("expensive.Result", "NewArt")

	// Read from jsonPath
	res = testJsonService.GetDataFromJsonPath("expensive.Result")
	log.Println(res)

}
