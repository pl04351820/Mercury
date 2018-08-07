package JPath

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/oliveagle/jsonpath"
)


// TODO: Delete this file later.
func JPathRun() {

	data := `
	{
    "store": {
        "book": [
            {
                "category": "reference",
                "author": "Nigel Rees",
                "title": "Sayings of the Century",
                "price": 8.95
            },
            {
                "category": "fiction",
                "author": "Evelyn Waugh",
                "title": "Sword of Honour",
                "price": 12.99
            },
            {
                "category": "fiction",
                "author": "Herman Melville",
                "title": "Moby Dick",
                "isbn": "0-553-21311-3",
                "price": 8.99
            },
            {
                "category": "fiction",
                "author": "J. R. R. Tolkien",
                "title": "The Lord of the Rings",
                "isbn": "0-395-19395-8",
                "price": 22.99
            }
        ],
        "bicycle": {
            "color": "red",
            "price": 19.95
        }
    },
    "expensive": "10"
	}
	`
	// JsonPath
	var json_data interface{}
	json.Unmarshal([]byte(data), &json_data)
	fmt.Printf("The data before json %+v \n", json_data)
	res, err := jsonpath.JsonPathLookup(json_data, "$.store.book[0].category")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(res)

	// Gabs
	jsonParsed, _ := gabs.ParseJSON([]byte(data))

	value, _ := jsonParsed.Path("expensive").Data().(string)
	fmt.Println(value)
	jsonParsed.SetP("datee", "expensive")
	value, _ = jsonParsed.Path("expensive").Data().(string)
	fmt.Println(value)
}
