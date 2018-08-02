package States

import (
	"Mercury/pkg/Type"
	"strconv"
	"log"
	"time"
)

type Wait struct {
	Common CommonField
	WaitType string
	WaitValue interface{}
	Next string
	End bool
}

func (w *Wait) run(events []byte) (string, []byte){
	return w.Next, events
}



// Legacy

func WaitState(task Type.Task, events []byte) (string, []byte) {
	// Only Support second for now.
	// Todo: Implement timestamp
	if task.Seconds != "" {
		timeDuration, err := strconv.Atoi(task.Seconds)
		if err != nil {
			log.Printf("Cannot transfer string to integer", err.Error())
		}

		time.Sleep(time.Duration(timeDuration) * time.Second)
	}
	return task.Next, events
}
