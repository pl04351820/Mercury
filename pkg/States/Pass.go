package States

type Pass struct {
	Common     CommonField
	Next       string `json:"Next"`
	End        bool   `json:"End"`
	ResultPath string `json:"ResultPath"`
	InputPath  string `json:"InputPath"`
	OutputPath string `json:"OutputPath"`
}

func (p *Pass) run() string{
	// To be implement check
	return ""
}
