package States

type CommonField struct {
	Comment string `json:"Comment"`
	Type string `json:"Type"`
	Error interface{} `json:"Error"`
	Cause interface{} `json:"Cause"`
	InputPath string `json:"InputPath"`
	OutputPath string `json:"OutputPath"`
	Events interface{}
}