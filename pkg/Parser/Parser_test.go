package Parser

import (
	"testing"
	"fmt"
)

func TestParser(t *testing.T) {
	res := Parser("./demo.json")
	fmt.Printf("%+v\n", res)
}

