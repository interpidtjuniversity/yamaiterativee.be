package db

import (
	"fmt"
	"strings"
	"xorm.io/builder"
)

type ServerType int

const (
	DEV ServerType = iota
	STABLE
	TEST
	PRE
	PROD
	UNKNOWN
)

func (st ServerType) FromInt(t int) ServerType {
	switch t {
	case 0:
		return DEV
	case 1:
		return STABLE
	case 2:
		return TEST
	case 3:
		return PRE
	case 4:
		return PROD
	default:
		return UNKNOWN
	}
}

func (st ServerType) ToString() string {
	switch st {
	case DEV:
		return "dev"
	case STABLE:
		return "stable"
	case TEST:
		return "test"
	case PRE:
		return "pre"
	case PROD:
		return "prod"
	default:
		return "unknown"
	}
}

func (st ServerType) FromString(str string) ServerType {
	switch str {
	case "dev":
		return DEV
	case "stable":
		return STABLE
	case "test":
		return TEST
	case "pre":
		return PRE
	case "prod":
		return PROD
	default:
		return st
	}
}

type ServerState int

const (
	APPLYING ServerState = iota
	DEPLOYING
	IDLE
	RUNNING
	STOPPED
)

func (ss ServerState) ToString() string{
	switch ss {
	case APPLYING:
		return "applying"
	case DEPLOYING:
		return "deploying"
	case IDLE:
		return "idle"
	case RUNNING:
		return "running"
	case STOPPED:
		return "stopped"
	default:
		return ""
	}
}

type Server struct {
	ID          int64       `xorm:"id autoincr pk"`
	Name        string      `xorm:"name"`
	AppOwner    string      `xorm:"app_owner"`
	AppName     string      `xorm:"app_name"`
	AppType     string      `xorm:"app_type"`
	IP          string      `xorm:"ip"`
	State       ServerState `xorm:"state"`
	Owner       string      `xorm:"owner"`
	Type        ServerType  `xorm:"type"`
	Branch      string      `xorm:"branch"`
	NetWork     string      `xorm:"net_work"`
	PortMapping string      `xorm:"port_mapping"`
	CreatedTime string      `xorm:"created_time"`
	IterationId int64       `xorm:"iteration_id"`
	DeployId    string      `xorm:"deploy_id"`
	GroupId     string      `xorm:"group_id"`
	ReleaseId   string      `xorm:"release_id"`
}

func InsertServer(server *Server) (bool, error) {
	exist, err := x.Table("server").Cols("name").Where(builder.Eq{"name": server.Name}).Get(new(Server))
	if exist || err!=nil {
		return false, err
	}
	_, err = x.Table("server").Insert(server)
	if err!=nil {
		return false, err
	}
	return true, nil
}

func GetServerByName(name string) (*Server, error) {
	server := new(Server)
	exist,err := x.Table("server").Where(builder.Eq{"name": name}).Get(server)
	if !exist || err!=nil {
		return nil, err
	}
	return server, nil
}

func DeleteServer(name string) (bool, error) {
	_,err := x.Table("server").Where(builder.Eq{"name": name}).Delete(new(Server))
	if err!=nil {
		return false, err
	}
	return true, nil
}

func UpdateServerAfterApply(name, ip string) (bool, error) {
	server := &Server{IP: ip, State: IDLE}
	_, err := x.Table("server").Where(builder.Eq{"name": name}).Update(server)
	if err!=nil {
		return false, err
	}
	return true,nil
}

func UpdateServerAfterDeploy(name string) (bool, error) {
	server := &Server{State: RUNNING}
	_, err := x.Table("server").Where(builder.Eq{"name": name}).Update(server)
	if err!=nil {
		return false, err
	}
	return true,nil
}

func GetServerByOwnerName(name string) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Where(builder.Eq{"owner": name}).Find(&servers)
	return servers, err
}

func UpdateServerDeployId(name, deployId string, state ServerState) (bool, error) {
	server := &Server{DeployId: deployId, State: state}
	_, err := x.Table("server").Cols("state", "deploy_id").Where(builder.Eq{"name": name}).Update(server)
	if err!=nil {
		return false, err
	}
	return true,nil
}

func GetDeployIdByServerName(name string) (string, error) {
	server := new(Server)
	exist, err := x.Table("server").Cols("deploy_id").Where(builder.Eq{"name":name}).Get(server)
	if err!=nil || !exist {
		return "", err
	}
	return server.DeployId, nil
}

func GetServerTypeAndGroupByIP(ip string) (string, string, error) {
	server := new(Server)
	exist, err := x.Table("server").Cols("name", "group_id").Where(builder.Eq{"ip":ip}).Get(server)
	if err!=nil || !exist {
		return "", "", err
	}
	array := strings.Split(server.Name, ".")
	if len(array) != 3{
		return "", "", err
	}
	return array[len(array)-1], server.GroupId, nil
}

func GetSameGroupServerByGroupIdAndServiceName(groupId string, serviceName string) ([]*Server, error) {
	var servers []*Server
	serviceName = strings.ReplaceAll(serviceName, "-", ".")
	err := x.Table("server").Cols("name", "ip").Where(builder.Eq{"group_id":groupId}.And(builder.Eq{"state": RUNNING}).And(builder.Like{"name", fmt.Sprintf("%s%s%s", "%", serviceName, "%")})).Find(&servers)
	return servers, err
}

func GetDevServerByAppOwnerAndName(appOwner, appName string) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Cols("name").Where(builder.Eq{"app_owner":appOwner, "app_name":appName, "type":DEV, "group_id":""}).Find(&servers)
	return servers, err
}

func GetGroupedDevServerByIterId(iterId int64) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Cols("name","owner","state","group_id").Where(builder.Eq{"iteration_id":iterId, "type":DEV}.And(builder.Neq{"group_id":""})).Find(&servers)
	return servers, err
}

func BranchUpdateServerGroup(names []string, groupId string) error {
	server := Server{GroupId: groupId}
	_, err := x.Table("server").Cols("group_id").Where(builder.In("name", names)).Update(&server)
	return err
}

func BranchQueryServerByGroupId(groupIds []string) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Cols("name","app_owner","app_name","owner","state","group_id").Where(builder.In("group_id", groupIds)).Find(&servers)
	return servers, err
}

func BranchQueryServerNameByGroupId(groupIds []string) ([]string, error) {
	var servers []*Server
	var serverNames []string
	err := x.Table("server").Cols("name").Where(builder.In("group_id", groupIds)).Find(&servers)
	for _, v := range servers {
		serverNames = append(serverNames, v.Name)
	}
	return serverNames, err
}

func GetServerIPByServerName(serverName string) string {
	server := new(Server)
	x.Table("server").Cols("ip").Where(builder.Eq{"name":serverName}).Get(server)
	return server.IP
}

func GetServerByAppAndOwner(appOwner, appName, owner string) ([]*Server, error){
	var servers []*Server
	err := x.Table("server").Where(builder.Eq{"app_owner":appOwner, "app_name":appName, "owner":owner}).Find(&servers)
	return servers, err
}

func GetServerByType(serverType ServerType) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Where(builder.Neq{"name":"", "ip":""}.And(builder.Eq{"type":serverType})).Find(&servers)
	return servers, err
}

func GetApplicationProdServer(appOwner, appName string) ([]*Server, error) {
	var servers []*Server
	err := x.Table("server").Cols("name", "ip").Where(builder.Eq{"app_owner":appOwner, "app_name":appName,"type":PROD}).Find(&servers)
	return servers, err
}

func GetServerStateAndTypeByName(name string) *Server {
	server := new(Server)
	x.Table("server").Cols("state", "type").Where(builder.Eq{"name":name}).Get(server)
	return server
}

func UpdateProdServerReleaseId(appOwner, appName, releaseId string) error {
	if releaseId == "" {
		return fmt.Errorf("release_id can not be empty")
	}
	server := &Server{ReleaseId: releaseId}
	_, err := x.Table("server").Cols("release_id").Where(builder.Eq{"app_owner":appOwner, "app_name": appName, "type":PROD}).Update(server)
	return err
}