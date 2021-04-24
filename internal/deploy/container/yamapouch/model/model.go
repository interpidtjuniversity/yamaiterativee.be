package model

type ContainerInfo struct {
	Pid         string   `json:"pid"`
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Command     string   `json:"command"`
	CreatedTime string   `json:"createTime"`
	Status      string   `json:"status"`
	Volume      string   `json:"volume"`
	PortMapping []string `json:"portmapping"`
	Ip          string   `json:"ip"`
	NetWorkName string   `json:"networkname"`
}

type ContainerCreateInfo struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}