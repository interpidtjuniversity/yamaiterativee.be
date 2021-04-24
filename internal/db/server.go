package db


type Server struct {
	ID          int64  `xorm:"id autoincr pk"`
	Name        string `xorm:"name"`
	AppId       int64  `xorm:"app_id"`
	IP          string `xorm:"ip"`
	State       string `xorm:"state"`
	Owner       string `xorm:"owner"`
	Type        string `xorm:"type"`
	Branch      string `xorm:"branch"`
	NetWork     string `xorm:"net_work"`
	PortMapping string `xorm:"port_mapping"`
	CreatedTime string `xorm:"created_time"`
}

func NewServer() {

}

