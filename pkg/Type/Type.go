package Type

type StateTransit struct{
	Variable interface{} `json:"Variable"`
	NumericEquals interface{} `json:"NumericEquals"`
	Next string `json:"Next"`
}

type Task struct{
	Type interface{}	`json:"Type"`
	Resource string	`json:"Resource"`
	Choices []StateTransit `json:"Choices"`
	Next string `json:"Next"`
	End bool `json:"End"`
	Error interface{} `json:"Error"`
	Cause interface{} `json:"Cause"`
}

type Job struct {
	Comment interface{} `json:"Comment"`
	StartAt string `json:"StartAt"`
	States map[string]Task `json:"States"`
}