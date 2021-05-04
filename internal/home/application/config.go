package application

import (
	"encoding/json"
	"net"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/resource"
)

var (
	JavaSpringConfig db.Config
	PythonConfig db.Config
	GolangConfig db.Config
)

type JavaSpringDynamicConfig struct {
	CONSUL_HOST string
	CONSUL_PORT string
	ZIPKIN_URL  string
	HOST_NAME   string
	HOST_TAGs   string
	INSTANCE_ID string
	APP_NAME    string
}

var JAVA_SPRING_DYNAMIC_CONFIG = JavaSpringDynamicConfig{
	CONSUL_HOST: "spring.cloud.consul.host",
	CONSUL_PORT: "spring.cloud.consul.port",
	ZIPKIN_URL:  "spring.zipkin.base-url",
	HOST_NAME:   "spring.cloud.consul.discovery.hostname",
	HOST_TAGs:   "spring.cloud.consul.discovery.tags",
	INSTANCE_ID: "spring.cloud.consul.discovery.instance-id",
	APP_NAME:    "spring.application.name",
}

func InitConfig() {
	initJavaSpringConfig()
}

func initJavaSpringConfig() {
	JavaSpringConfig = db.GetJavaSpringConfig()
	// consul use yamaiterativee proxy
	JavaSpringConfig.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.CONSUL_HOST, getLocalIPv4Address())
	JavaSpringConfig.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.CONSUL_PORT, 4000)
	// zipkin use global zipkin
	JavaSpringConfig.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.ZIPKIN_URL, resource.GLOBAL_ZIPKIN_IP)
	// mysql use global mysql
	JavaSpringConfig.SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.ZIPKIN_URL, resource.GLOBAL_MYSQL_IP)
}

func GetJavaSpringConfig() ([]byte, error) {
	var configItems []db.ConfigItem
	for _, v := range JavaSpringConfig.ConfigItems {
		if v.Displayable {
			configItems = append(configItems, v)
		}
	}
	data, err := json.Marshal(configItems)
	return data, err
}

func getLocalIPv4Address() (ipv4Address string) {
	//获取所有网卡
	addrs, _ := net.InterfaceAddrs()
	//遍历
	for _, addr := range addrs {
		//取网络地址的网卡的信息
		ipNet, isIpNet := addr.(*net.IPNet)
		//是网卡并且不是本地环回网卡
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			//能正常转成ipv4
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}
	return
}