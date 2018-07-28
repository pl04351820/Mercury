package States

import (
	"Mercury/pkg/Type"
	"fmt"
	"strconv"
	"time"
)

/*
The rule of Input and output processing.

@Common Field.
	Type
	Next
	End
	Comment
	InputPath
	OutputPath

@Input
	State Field
	Pass Transition (Json)
@Output
	Next Field or End State
	Pass Transition (Json)
*/

func WaitState(task Type.Task, events []byte) (string, []byte) {
	// Only Support second for now.
	// Todo: Implement timestamp
	if task.Seconds != "" {
		timeDuration, err := strconv.Atoi(task.Seconds)
		if err != nil {
			fmt.Printf("Cannot transfer string to integer", err.Error())
		}

		time.Sleep(time.Duration(timeDuration) * time.Second)
	}
	return task.Next, events
}
