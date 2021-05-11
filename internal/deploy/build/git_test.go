package build

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClone(t *testing.T) {
	_, _, err := Clone("http://localhost:3002/interpidtjuniversity/miniselfop", "master","dasfffewf", "miniselfop")
	assert.Nil(t, err)
}
