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


type Server struct {
	ID          int64      `xorm:"id autoincr pk"`
	Name        string     `xorm:"name"`
	AppOwner    string     `xorm:"app_owner"`
	AppName     string     `xorm:"app_name"`
	AppType     string     `xorm:"app_type"`
	IP          string     `xorm:"ip"`
	State       string     `xorm:"state"`
	Owner       string     `xorm:"owner"`
	Type        ServerType `xorm:"type"`
	Branch      string     `xorm:"branch"`
	NetWork     string     `xorm:"net_work"`
	PortMapping string     `xorm:"port_mapping"`
	CreatedTime string     `xorm:"created_time"`
	IterationId int64      `xorm:"iteration_id"`
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

