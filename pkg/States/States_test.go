package States

import (
	"testing"
)

func TestState(t *testing.T) {
	// Test Wait State
	commonElement := CommonField{Comment: "Comment", Type: "Wait", Events: []byte(`{"expensive":{"Result":"Art", "FirstChoice":3}}`)}
	WaitState := Wait{Common: commonElement, WaitType: "Seconds", WaitValue: 1, Next: "edit"}
	WaitState.Run()

	// Test Pass State
	PassState := Pass{Common: commonElement, Next: "edit"}
	PassState.Run()

	// Test Task State
	TaskState := Task{Common: commonElement, Next: "edit", InputPath: "", ResultPath: "", OutputPath: ""}
	TaskState.Run()

	// Test Choice State
	ChoiceState := Choice{Common: commonElement, Choices: []StateTransition{
		{Next: "edit", OperationType: "NumericGreaterThan", OperationBase: 5.0, OperationRef: "expensive.FirstChoice"},
		{Next: "edit2", OperationType: "NumericLessThan", OperationBase: 5.0, OperationRef: "expensive.FirstChoice"}}}
	ChoiceState.Run()
}
