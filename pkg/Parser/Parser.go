package Parser

import (
	"Mercury/pkg/Path"
)


type Parser struct{
	PathSvc Path.JsonPathService
}

func NewParser(RawData []byte) (Parser){
	PathServer := Path.NewJsonPathService(RawData)
	newParser := Parser{PathSvc:PathServer}
	return newParser
}

func (p *Parser) ParseTask(stateName string) interface{}{
	// return specific type
	stateType := p.PathSvc.GetDataFromJsonPath(".States" + "." + stateName + "." + "Type")

	// TODO: Parse data according to different data type.


	return 	stateType
}
