package server

import (
	"fmt"
	"testing"
)

func TestFlush(t *testing.T) {
	var v interface{}
	v=true
	fmt.Print(fmt.Sprintf("%s", v))
	v=100
	fmt.Print(fmt.Sprintf("%s", v))
	v="Dasda"
	fmt.Print(fmt.Sprintf("%s", v))
}

