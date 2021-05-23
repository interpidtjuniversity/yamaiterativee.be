package bean

import (
	"fmt"
)

type BranchCompileBean struct {
	CompileBean
}

// 0     1       2       3
//repo branch execPath logPath
func (bcb *BranchCompileBean) Execute(stringArgs []string, env *map[string]interface{}) error {
	// 1.parse arg
	if len(stringArgs) != 4 {
		return fmt.Errorf("arguement error")
	}
	err := bcb.clone(stringArgs[0], stringArgs[1], stringArgs[2], stringArgs[3])

	return err
}
