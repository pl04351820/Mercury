package Parser

import (
	"Mercury/pkg/Type"
	"io/ioutil"
	"encoding/json"
)

func Parser(template string) Type.Job{
	var job Type.Job
	file, _ := ioutil.ReadFile(template)
	json.Unmarshal(file, &job)
	return job
}