package Type

type StateTransit struct{
	Variable string
	NumericEquals string
	Next Task
}

type Task struct{
	Type string
	Resource string
	Choices StateTransit
	Next Task
	End string
	Error string
	Cause string
}

type Job struct {
	
}