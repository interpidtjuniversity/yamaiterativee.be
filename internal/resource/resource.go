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
	GLOBAL_ZIPKIN_IP string
	GLOBAL_MYSQL = "GLOBAL_MYSQL"
	GLOBAL_CONSUL = "GLOBAL_CONSUL"
	GLOBAL_ZIPKIN = "GLOBAL_ZIPKIN"
)

func InitResource() {
	network.InitNetwork()
	mysql,_ := db.GetResourceByName(GLOBAL_MYSQL)
	if mysql==nil {
		mysqlIP, err := server.NewMysqlServer(GLOBAL_MYSQL, network.NAME, ENV)
		if err != nil {
			panic("error while create global mysql")
		}
		GLOBAL_MYSQL_IP = mysqlIP
		db.InsertResource(&db.Resource{
			Name: GLOBAL_MYSQL,
			Value: mysqlIP,
		})
	} else{
		GLOBAL_MYSQL_IP = mysql.Value
	}
	consul,_ := db.GetResourceByName(GLOBAL_CONSUL)
	if consul==nil {
		consulIP, err := server.NewConsulServer(GLOBAL_CONSUL, network.NAME, ENV)
		if err != nil {
			panic("error while create global consul")
		}
		GLOBAL_CONSUL_IP = consulIP
		db.InsertResource(&db.Resource{
			Name: GLOBAL_CONSUL,
			Value: consulIP,
		})
	} else {
		GLOBAL_CONSUL_IP = consul.Value
	}
	zipkin,_ := db.GetResourceByName(GLOBAL_ZIPKIN)
	if zipkin==nil {
		zipkinIP, err := server.NewZipkinServer(GLOBAL_ZIPKIN, network.NAME, ENV)
		if err != nil {
			panic("error while create global zipkin")
		}
		GLOBAL_ZIPKIN_IP = zipkinIP
		db.InsertResource(&db.Resource{
			Name: GLOBAL_ZIPKIN,
			Value: zipkinIP,
		})
	} else {
		GLOBAL_ZIPKIN_IP = zipkin.Value
	}
}
