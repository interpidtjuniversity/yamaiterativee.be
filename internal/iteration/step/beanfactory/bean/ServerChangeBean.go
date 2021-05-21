package bean

import (
	"fmt"
	"strconv"
	"time"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/util"
)

type ServerChangeBean struct {
	Bean
}
//    0       1        2       3         4         5
// appOwner appName   owner  appType   iterId   serverType
func (sb *ServerChangeBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	if len(stringArgs) != 6 {
		return fmt.Errorf("arguement error")
	}
	iterId, _ := strconv.Atoi(stringArgs[4])
	st, _ := strconv.Atoi(stringArgs[5])
	serverType := db.ServerType(st)
	return sb.change(stringArgs[0], stringArgs[1], stringArgs[2], stringArgs[3], int64(iterId), serverType, env)
}

func (sb *ServerChangeBean) change(appOwner, appName, owner, appType string, iterId int64, serverType db.ServerType, env *map[string]interface{}) error{

	newServer := &db.Server{AppOwner: appOwner, AppName: appName, IterationId: iterId, Owner: owner, AppType: appType,
		Type: serverType, State: db.APPLYING, CreatedTime: time.Now().Format("2006-01-01 15:04:05"),
		Name: util.GenerateRandomStringWithSuffix(20, fmt.Sprintf("%s.%s.%s", appOwner, appName, serverType.ToString())),
	}

	networkName, _ := db.GetApplicationNetworkByOwnerAndRepo(appOwner, appName)
	newServer.NetWork = networkName
	_, err := db.InsertServer(newServer)
	if err != nil {
		return fmt.Errorf("error while create server, error: %s", err)
	}
	(*env)["networkName"] = networkName
	(*env)["serverName"] = newServer.Name
	return nil
}

