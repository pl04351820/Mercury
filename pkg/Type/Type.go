package Type

type StateTransit struct {
	Variable           string `json:"Variable"`
	NumericEquals      int    `json:"NumericEquals"`
	NumericGreaterThan int    `json:"NumericGreaterThan"`
	NumericLessThan    int    `json:"NumericLessThan"`
	Next               string `json:"Next"`
}

type Task struct {
	Type     interface{}    `json:"Type"`
	Resource string         `json:"Resource"`
	Choices  []StateTransit `json:"Choices"`
	Next     string         `json:"Next"`
	End      bool           `json:"End"`
	Error    interface{}    `json:"Error"`
	Cause    interface{}    `json:"Cause"`
	Seconds  string         `json:"Seconds"`
	Default  string         `json:"Default"`
	Status   bool
}

type Job struct {
	Comment interface{}     `json:"Comment"`
	StartAt string          `json:"StartAt"`
	States  map[string]Task `json:"States"`
}

type ESType struct {
	TaskName string
	LogInfo  string
}

type JobState struct {
	JobName    string
	StatueInfo map[string]bool
}
