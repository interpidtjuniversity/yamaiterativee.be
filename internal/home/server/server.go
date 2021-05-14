package server

import (
	"encoding/json"
	"fmt"
	"time"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/deploy/container/yamapouch/command"
	"yama.io/yamaIterativeE/internal/util"
)

const (
	// root-miniselfop-cl5664dwed-dev
	// system-mysql-zcnelw05-dev
	// userName-appName-id-env
	SERVER_NAME = "%s-%s-%s-%s"
)

type ServerData struct {
	Name         string `json:"name"`
	AppOwner     string `json:"appOwner"`
	AppName      string `json:"appName"`
	IP           string `json:"ip"`
	DeployBranch string `json:"deployBranch"`
	Env          string `json:"env"`
	Owner        string `json:"owner"`
	State        string `json:"state"`
	ApplyTime    string `json:"applyTime"`
	IterId       int64  `json:"iterId"`
}

type GroupServerData struct {
	// front use
	Rdm            string `json:"rdm"`
	AppOwner       string `json:"appOwner"`
	AppName        string `json:"appName"`
	AppServer      string `json:"appServer"`
	// back use
	AppServerOwner string `json:"appServerOwner"`
	AppServerState int    `json:"appServerState"`
}

func NewServer(c *context.Context) []byte {
	appType := c.Query("appType")
	iterBranch := c.Query("iterBranch")
	appOwner := c.Query("appOwner")
	appName := c.Query("appName")
	iterId := c.QueryInt64("iterId")
	owner := c.Query("owner")
	serverType := db.ServerType(c.QueryInt("serverType"))

	newServer := &db.Server{AppOwner: appOwner, AppName: appName, IterationId: iterId, Owner: owner, AppType: appType,
		Branch: iterBranch, Type: serverType, State: db.APPLYING, CreatedTime: time.Now().Format("2006-01-01 15:04:05"),
		Name: util.GenerateRandomStringWithSuffix(20, fmt.Sprintf("%s.%s.%s", appOwner, appName, serverType.ToString())),
	}

	networkName, _ := db.GetApplicationNetworkByOwnerAndRepo(appOwner, appName)
	newServer.NetWork = networkName
	_, err := db.InsertServer(newServer)
	if err != nil {
		return []byte(fmt.Errorf("error while create server, error: %s", err).Error())
	}

	// 1. invoke New...Server by serverType
	var ip string
	switch appType {
	case db.JAVA_SPRING:
		ip,_ = NewJavaApplicationServer(newServer.Name, networkName)
		break
	default:
		return []byte(fmt.Errorf("error while create server, error: unsupport appType %s", appType).Error())
	}
	// 2. update server ip
	db.UpdateServerAfterApply(newServer.Name, ip)

	return nil
}


func GetUserAllServers(c *context.Context) []byte {
	owner := c.Params(":username")
	servers, _ := db.GetServerByOwnerName(owner)
	var serverDatas []ServerData
	for _, server := range servers {
		data := ServerData{
			Name: server.Name,
			AppOwner: server.AppOwner,
			AppName: server.AppName,
			IP: server.IP,
			DeployBranch: server.Branch,
			Env: server.Type.ToString(),
			State: server.State.ToString(),
			ApplyTime: server.CreatedTime,
			IterId: server.IterationId,
			Owner: owner,
		}
		serverDatas = append(serverDatas, data)
	}
	data, _ := json.Marshal(serverDatas)
	return data
}


func GetAppDevServer(c *context.Context) []byte {
	appOwner := c.Params(":appOwner")
	appName := c.Params(":appName")

	servers, err := db.GetDevServerByAppOwnerAndName(appOwner, appName)
	if err != nil {
		return []byte("[]")
	}
	names := make([]string, 0)
	for _, server := range servers {
		names = append(names, server.Name)
	}
	data, _ := json.Marshal(names)

	return data
}

func CreateIterationDebugGroup(c *context.Context) []byte {
	groupServerBytes := []byte(c.Query("groupServer"))
	groupServer := make([]GroupServerData, 1)
	json.Unmarshal(groupServerBytes, &groupServer)
	var names []string
	for _, gs := range groupServer {
		if gs.AppServer!="" {
			names = append(names, gs.AppServer)
		}
	}
	groupId := util.GenerateRandomStringWithSuffix(20,"")
	if err := db.BranchUpdateServerGroup(names, groupId); err!=nil {
		return []byte(fmt.Sprintf("error while bind group, err: %v", err))
	}

	return []byte("success")
}

func QueryIterationDebugGroup(c *context.Context) []byte {
	iterId := c.ParamsInt64("iterId")
	servers, _ := db.GetGroupedDevServerByIterId(iterId)
	var groupIds []string
	for _, server := range servers {
		groupIds = append(groupIds, server.GroupId)
	}
	allServers, _ := db.BranchQueryServerByGroupId(groupIds)

	groupServerMap := make(map[string][]GroupServerData)
	for _, server := range allServers {
		groupServerMap[server.GroupId] = append(groupServerMap[server.GroupId], GroupServerData{
			AppServer: server.Name,
			AppServerOwner: server.Owner,
			AppServerState: int(server.State),
			AppOwner: server.AppOwner,
			AppName: server.AppName,
			Rdm: util.GenerateRandomStringWithSuffix(10,""),
		})
	}
	data, _:= json.Marshal(groupServerMap)

	return data
}

func DeleteServerInIterationDebugGroup(c *context.Context) []byte {
	return nil
}

func NewJavaApplicationServer(containerName, network string) (string, error){
	time.Sleep(time.Duration(5)*time.Second)
	ip, err := command.CreateContainer(containerName, network, "", "JavaImage", "top -b")
	return ip, err
}

func NewMysqlServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "mysql", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "MysqlImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, "/bin/chmod 666 /dev/null#/usr/sbin/service mysql start")
	return ip, err
}

func NewConsulServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "consul", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "ConsulImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, fmt.Sprintf("/bin/chmod 777 standalone-consul.sh#./standalone-consul.sh %s %s %s", containerName, ip, containerName))
	return ip, err
}

func NewZipkinServer(containerName, network, env string) (string, error){
	if containerName == "" {
		containerName = fmt.Sprintf(SERVER_NAME, "system", "zipkin", util.GenerateRandomStringWithSuffix(10, ""), env)
	}
	ip, err := command.CreateContainer(containerName, network, "", "ZipkinImage", "top -b")
	if err != nil || ip == ""{
		return ip, err
	}
	err = command.ExecuteBatchCommandOnceInContainer(containerName, "/bin/chmod 777 standalone-zipkin.sh#./standalone-zipkin.sh")
	return ip, err
}
