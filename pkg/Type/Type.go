package Type

// Use Reflection mechanism to check the type of structure.
type Job struct {
	JobName string
	Comment string                 `json:"Comment"`
	StartAt string                 `json:"StartAt"`
	States  map[string]interface{} `json:"States"`
	Events  string
}
