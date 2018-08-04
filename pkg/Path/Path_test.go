package Path

import (
	"log"
	"testing"
)

func TestNewJsonPathService(t *testing.T) {
	testJsonService := NewJsonPathService([]byte(`{"expensive":{"Result":"Art"}}`))

	// Read from jsonPath
	res := testJsonService.GetDataFromJsonPath("expensive.Result")
	log.Println(res)

	// Retrieve to jsonPath
	testJsonService.WriteDataToJsonPath("expensive.Result1", "NewArt")

	// Read from jsonPath
	res = testJsonService.GetDataFromJsonPath("expensive.Result1")
	log.Println(res)

}
