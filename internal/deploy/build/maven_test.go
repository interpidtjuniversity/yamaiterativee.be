package build

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMavenCompile(t *testing.T) {
	_, err := MavenCompile("/var/run/yama/yamaIterativeE/dasfffewf/miniselfop")
	assert.Nil(t, err)
}

func TestMavenInstall(t *testing.T) {
	_,err := MavenInstall("/var/run/yama/yamaIterativeE/dasfffewf/miniselfop")
	assert.Nil(t, err)
}