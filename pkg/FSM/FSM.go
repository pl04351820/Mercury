package FSM

import (
	"Mercury/pkg/Parser"
	"fmt"
	"Mercury/pkg/Type"
	"github.com/aws/aws-sdk-go"
)

func ExecuteTask(input string, output string, task Type.Task)string{
	// Invoke Lambda function with events here.


	// Return next state
	if task.End == True{
		return "End_Signal"
	}else{
		return task.Next
	}

}

func FSM(){
	FSM := Parser.Parser("./demo.json")
	fmt.Printf("%+v\n", FSM)

	// Start to execute FSM.
	StateName := FSM.StartAt
	fmt.Println(StateName)

	EndFlag := false
	for EndFlag == false{
		NextState := ExecuteTask(input, ouput, FSM.States[StateName])
		if NextState == "End_Signal"{
			EndFlag = true
		}
	}

}