package db

type ServerType int

const (
	DEV ServerType = iota
	STABLE
	TEST
	PRE
	PROD
)

type Server struct {
	ID          int64      `xorm:"id autoincr pk"`
	Name        string     `xorm:"name"`
	AppId       int64      `xorm:"app_id"`
	IP          string     `xorm:"ip"`
	State       string     `xorm:"state"`
	Owner       string     `xorm:"owner"`
	Type        ServerType `xorm:"type"`
	Branch      string     `xorm:"branch"`
	NetWork     string     `xorm:"net_work"`
	PortMapping string     `xorm:"port_mapping"`
	CreatedTime string     `xorm:"created_time"`
}

func NewServer() {

}

