package States

import (
	"time"
)

type Wait struct {
	Common     CommonField
	WaitType   string
	WaitValue  interface{}
	Next       string
	End        bool
	InputPath  string `json:"InputPath"`
	OutputPath string `json:"OutputPath"`
}

func (w *Wait) Run() string {
	switch w.WaitType {
	case "Seconds":
		time.Sleep(time.Duration(w.WaitValue.(int)) * time.Second)
		break
	}
	return ""
}
