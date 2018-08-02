package States

type Pass struct {
	Common CommonField
	Next string `json:"Next"`
	End bool `json:"End"`
	ResultPath string `json:"ResultPath"`
}