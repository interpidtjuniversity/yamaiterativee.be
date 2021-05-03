package resource

import (
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/home/server"
)
import "yama.io/yamaIterativeE/internal/home/network"

const (
	ENV = "GLOBAL"
)

var(
	GLOBAL_MYSQL_IP string
	GLOBAL_CONSUL_IP string
	GLOAB_ZIPKIN_IP string
	GLOBAL_MYSQL = "GLOBAL_MYSQL"
	GLOBAL_CONSUL = "GLOBAL_CONSUL"
	GLOBAL_ZIPKIN = "GLOBAL_ZIPKIN"
)

func InitResource() {
	network.InitNetwork()
	if mysql,_ := db.GetResourceByName(GLOBAL_MYSQL); mysql==nil {
		mysqlIP, err := server.NewMysqlServer(GLOBAL_MYSQL, network.NAME, ENV)
		if err != nil {
			panic("error while create global mysql")
		}
		GLOBAL_MYSQL_IP = mysqlIP
	}
	if consul,_ := db.GetResourceByName(GLOBAL_CONSUL); consul==nil {
		consulIP, err := server.NewConsulServer(GLOBAL_CONSUL, network.NAME, ENV)
		if err != nil {
			panic("error while create global consul")
		}
		GLOBAL_CONSUL_IP = consulIP
	}
	if zipkin,_ := db.GetResourceByName(GLOBAL_ZIPKIN); zipkin==nil {
		zipkinIP, err := server.NewZipkinServer(GLOBAL_ZIPKIN, network.NAME, ENV)
		if err != nil {
			panic("error while create global zipkin")
		}
		GLOAB_ZIPKIN_IP = zipkinIP
	}
}
