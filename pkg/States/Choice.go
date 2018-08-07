package States

import (
	"Mercury/pkg/Path"
)

/*
If OperationType == "and" || "or":
	OperationValue is []StateTransition
else:
	OperationValue is SimpleElement
*/

/*
The type to be implemented in the future work:
	TODO: Implement the following type in future
	And
	Or
	TimestampEquals
	TimestampGreaterThan
	TimestampGreaterThanEquals
	TimestampLessThan
	TimestampLessThanEquals
*/
var CompareOperation = []interface{}{"BooleanEquals", "NumericEquals", "StringEquals", "NumericGreaterThan",
	"StringGreaterThan", "NumericGreaterThanEquals", "StringGreaterThanEquals",
	"NumericLessThan", "NumericLessThan", "StringLessThan", "NumericLessThanEquals",
	"StringLessThanEquals", "Not"}

type StateTransition struct {
	Next          string
	OperationType string
	OperationBase interface{}
	OperationRef  string
}

type Choice struct {
	Common     CommonField
	Choices    []StateTransition
	Default    string
	InputPath  string
	OutputPath string
	Next       string // Use this field to represent final next in dynamic.
}

func (c *Choice) run() string {
	c.Next = c.waterFall()
	return ""
}

// ForLoop, return the first one that return True
func (c *Choice) waterFall() string {
	pathService := Path.NewJsonPathService(c.Common.Events)
	var nextTask string

	for _, element := range c.Choices {
		operationValue := pathService.GetDataFromJsonPath(element.OperationRef)
		if c.oneCheck(element.OperationType, element.OperationBase, operationValue) {
			nextTask = element.Next
			return nextTask
		}
	}
	return c.Default
}

// One element check.
// Todo: Number Element Check here. Compatible for float and int.
func (c *Choice) oneCheck(operationType string, operationBase interface{}, operationValue interface{}) bool {
	switch operationType {
	case "BooleanEquals", "NumericEquals", "StringEquals":
		if operationValue == operationBase {
			return true
		}
	case "NumericGreaterThan":
		{
			if operationValue.(float64) > operationValue.(float64) {
				return true
			}
		}
	case "StringGreaterThan":
		{
			if operationValue.(string) > operationBase.(string) {
				return true
			}
		}
	case "NumericGreaterThanEquals":
		{
			if operationValue.(float64) >= operationBase.(float64) {
				return true
			}
		}
	case "StringGreaterThanEquals":
		{
			if operationValue.(string) >= operationBase.(string) {
				return true
			}
		}
	case "NumericLessThan":
		{
			if operationValue.(float64) < operationBase.(float64) {
				return true
			}
		}
	case "StringLessThan":
		{
			if operationValue.(string) < operationBase.(string) {
				return true
			}
		}
	case "NumericLessThanEquals":
		{
			if operationValue.(float64) <= operationBase.(float64) {
				return true
			}
		}
	case "StringLessThanEquals":
		{
			if operationValue.(string) <= operationBase.(string) {
				return true
			}
		}
	case "Not":
		if operationValue != operationBase {
			return true
		}
	}

	return false
}
