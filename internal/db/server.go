package db

import "xorm.io/builder"

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

type ServerState int

const (
	APPLYING ServerState = iota
	DEPLOYING
	IDLE
	RUNNING
	STOPPED
)

type Server struct {
	ID          int64       `xorm:"id autoincr pk"`
	Name        string      `xorm:"name"`
	AppOwner    string      `xorm:"app_owner"`
	AppName     string      `xorm:"app_name"`
	AppType     AppType     `xorm:"app_type"`
	IP          string      `xorm:"ip"`
	State       ServerState `xorm:"state"`
	Owner       string      `xorm:"owner"`
	Type        ServerType  `xorm:"type"`
	Branch      string      `xorm:"branch"`
	NetWork     string      `xorm:"net_work"`
	PortMapping string      `xorm:"port_mapping"`
	CreatedTime string      `xorm:"created_time"`
	IterationId int64       `xorm:"iteration_id"`
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

func ReleaseServer(name string) (bool, error) {
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

