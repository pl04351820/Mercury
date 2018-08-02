package Type

// Use

type Job struct {
	JobName string
	Comment string	        `json:"Comment"`
	StartAt string          `json:"StartAt"`
	States  map[string]Task `json:"States"`
}

type State struct {
	Type State
	Data interface{}
}