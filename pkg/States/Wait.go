package States

import (
	"time"
)

type Wait struct {
	WaitType string
	WaitValue interface{}
	Next string
	End bool
}

func (w *Wait) run()(string){
	switch w.WaitType {
	case "Seconds":
		time.Sleep(time.Duration(w.WaitValue.(int)) * time.Second)
	}
	return ""
}