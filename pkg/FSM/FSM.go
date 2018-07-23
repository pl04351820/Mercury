package FSM

import (
	"Mercury/pkg/Parser"
	"fmt"
	"Mercury/pkg/Type"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"encoding/json"
)

func ExecuteTask(payload []byte, task Type.Task)(string, []byte){
	// Invoke Lambda function with events here.
	ses := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := lambda.New(ses, &aws.Config{Region: aws.String("us-east-1")})

	result, err := svc.Invoke(&lambda.InvokeInput{FunctionName:aws.String(task.Resource),Payload:payload})
	if err != nil{
		fmt.Printf("Cannot Invoke because %v \n", err)
		os.Exit(0)
	}

	fmt.Printf("%v \n",string(result.Payload))

	// Return next state
	if task.End == true{
		return "End_Signal", result.Payload
	}else{
		return task.Next, result.Payload
	}

}

func FSM(){
	FSM := Parser.Parser("./demo.json")
	fmt.Printf("%+v\n", FSM)

	// Start to execute FSM.
	StateName := FSM.StartAt
	fmt.Println(StateName)

	// Init Input Events
	input := make(map[string]string)
	input["who"] = "led"
	payload, _ := json.Marshal(input)

	EndFlag := false

	//CurTask := FSM.States[StateName]
	for EndFlag == false{
		NextState, output := ExecuteTask(payload, FSM.States[StateName])

		if NextState == "End_Signal" {
			EndFlag = true
		}else{
			StateName = NextState
			payload = output
		}
	}
}