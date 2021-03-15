package invokerImpl

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PipeLIneService(t *testing.T) {
	response := InvokeRegisterMergeRequestService()
	assert.NotEqual(t, response.ShowDiffUri, "")
	fmt.Print(response.ShowDiffUri)
	fmt.Print("\n")
	fmt.Print(response.MRId)
}