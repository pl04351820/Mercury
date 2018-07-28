package States

import (
	"Mercury/pkg/Type"
	"encoding/json"
	"github.com/oliveagle/jsonpath"
	"log"
)

func ChoiceState(task Type.Task, events []byte) (string, []byte) {
	// JsonPaTh to read state.
	next := ""
	for _, element := range task.Choices {
		var value int
		json.Unmarshal(events, &value)
		res, err := jsonpath.JsonPathLookup(value, element.Variable)
		resData := res.(int)
		if err != nil {
			log.Println("Error Happen when loading Json Path")
		}
		if element.NumericGreaterThan != 0 {
			if resData > element.NumericGreaterThan {
				next = element.Next
			}
		} else if element.NumericLessThan != 0 {
			if resData < element.NumericLessThan {
				next = element.Next
			}
		} else if element.NumericEquals != 0 {
			if resData == element.NumericEquals {
				next = element.Next
			}
		}
	}
	return next, events
}
