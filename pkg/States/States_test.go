package States

import (
	"testing"
)

/*
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
 */


func TestState(t *testing.T) {
	// Test Wait State
	commonElement := CommonField{Comment:"Comment", Type:"Wait", Events: []byte(`{"expensive":{"Result":"Art"}}`)}
	WaitState := Wait{Common:commonElement, WaitType:"Seconds", WaitValue:1, Next:"edit"}
	WaitState.run()

	// Test Pass State
	PassState := Pass{Common:commonElement, Next:"edit"}
	PassState.run()

	// Test Task State
	TaskState := Task{Common:commonElement, Next:"edit", InputPath:"", ResultPath:"", OutputPath:""}
	TaskState.run()
}