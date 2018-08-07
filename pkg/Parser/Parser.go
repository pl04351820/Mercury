package Parser

import (
	"Mercury/pkg/Path"
	"Mercury/pkg/States"
	"github.com/deckarep/golang-set"
	"github.com/mitchellh/mapstructure"
	"reflect"
)

type Parser struct {
	PathSvc Path.JsonPathService
}

func NewParser(RawData []byte) Parser {
	PathServer := Path.NewJsonPathService(RawData)
	newParser := Parser{PathSvc: PathServer}
	return newParser
}

func (p *Parser) ParseTask(stateName string) interface{} {
	stateInfo := p.PathSvc.GetDataFromJsonPath("States" + "." + stateName)
	// TODO: Parse data according to different data type.
	commonField := States.CommonField{}
	mapstructure.Decode(stateInfo, &commonField)

	if commonField.Type == "Task" {
		taskState := States.Task{}
		mapstructure.Decode(stateInfo, &taskState)
		taskState.Common = commonField

		return taskState
	} else if commonField.Type == "Choice" {
		choiceState := States.Choice{}
		mapstructure.Decode(stateInfo, &choiceState)
		choiceState.Common = commonField
		choiceOption := mapset.NewSetFromSlice(States.CompareOperation)
		transitionsInfo := (p.PathSvc.GetDataFromJsonPath("States" + "." + stateName + "." + "Choices"))

		var choiceArr []States.StateTransition
		switch reflect.TypeOf(transitionsInfo).Kind() {
		case reflect.Slice:
			s := reflect.ValueOf(transitionsInfo)

			for i := 0; i < s.Len(); i++ {
				elementMap := s.Index(i).Interface().(map[string]interface{})
				choiceElement := States.StateTransition{}

				for k, v := range elementMap {
					if k == "Variable" {
						choiceElement.OperationRef = v.(string)
					} else if k == "Next" {
						choiceElement.Next = v.(string)
					} else if choiceOption.Contains(k) {
						choiceElement.OperationType = k
						choiceElement.OperationBase = v
					}
				}
				choiceArr = append(choiceArr, choiceElement)
			}
		}
		choiceState.Choices = choiceArr
		return choiceState

	} else if commonField.Type == "Pass" {
		passState := States.Pass{}
		mapstructure.Decode(stateInfo, &passState)
		passState.Common = commonField
		return passState
	}

	return ""
}
