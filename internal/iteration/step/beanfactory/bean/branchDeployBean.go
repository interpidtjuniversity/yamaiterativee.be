package bean

import (
	"encoding/json"
	"strconv"
)

type BranchDeployBean struct {
	CompileBean
	DeployBean
}

//    0       1      2      3         4        5             6
// appName execPath env serverName serverIP iterId(string) logPath
func (bdb *BranchDeployBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	var serverNames []string
	var serverIPs []string
	iterId, _ := strconv.Atoi(stringArgs[5])
	json.Unmarshal([]byte(stringArgs[3]), &serverNames)
	json.Unmarshal([]byte(stringArgs[4]), &serverIPs)

	for i:=0; i<len(serverNames); i++ {
		err := bdb.flushConfig(stringArgs[0], stringArgs[1], stringArgs[2], serverNames[i], serverIPs[i], int64(iterId))
		if err!=nil {
			return err
		}
		err = bdb.mavenInstall(stringArgs[0], stringArgs[1], stringArgs[6])
		if err!=nil {
			return err
		}
		err = bdb.deploy(stringArgs[0], stringArgs[1], serverNames[i], serverIPs[i])
		if err != nil {
			return err
		}
	}

	return nil
}
