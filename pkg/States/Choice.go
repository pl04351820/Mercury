package States

import (
	"encoding/json"
	"github.com/oliveagle/jsonpath"
	"log"
)

type StateTransition struct {
	Next           string
	OperationType  string
	OperationValue interface{}
}

type Choice struct {
	Common  CommonField
	Choices []StateTransition
	Default string
}

/*
The type to be implemented in the future work:
	And
	BooleanEquals
	Not
	NumericEquals
	NumericGreaterThan
	NumericGreaterThanEquals
	NumericLessThan
	NumericLessThanEquals
	Or
	StringEquals
	StringGreaterThan
	StringGreaterThanEquals
	StringLessThan
	StringLessThanEquals
	TimestampEquals
	TimestampGreaterThan
	TimestampGreaterThanEquals
	TimestampLessThan
	TimestampLessThanEquals
*/

// Legacy
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
