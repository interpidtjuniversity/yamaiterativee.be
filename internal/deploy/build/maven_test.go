package build

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMavenCompile(t *testing.T) {
	out, err := MavenCompile("/var/run/yama/yamaIterativeE/repo/ccf2007bb25b11e/miniselfop")
	assert.Nil(t, err)
	fmt.Print(string(out))
}

func TestMavenInstall(t *testing.T) {
	_,err := MavenInstall("/var/run/yama/yamaIterativeE/dasfffewf/miniselfop")
	assert.Nil(t, err)
}