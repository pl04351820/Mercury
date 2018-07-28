package Parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	res := Parser("./demo.json")
	fmt.Printf("%+v\n", res)
}
