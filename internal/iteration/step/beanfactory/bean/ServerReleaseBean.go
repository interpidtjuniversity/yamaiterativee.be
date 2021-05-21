package bean

import (
	"fmt"
	"yama.io/yamaIterativeE/internal/db"
)

type ServerReleaseBean struct {
	Bean
}

//
func (srb *ServerReleaseBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	if len(stringArgs)!=0 {
		return fmt.Errorf("arguement error")
	}
	return srb.release(stringArgs, env)
}

func (srb *ServerReleaseBean) release(stringArgs []string, env *map[string]interface{}) error {
	// 2. update server ip
	_, err := db.UpdateServerAfterApply((*env)["serverName"].(string), (*env)["serverIP"].(string))
	return err
}
