package config

import (
	"encoding/json"
	"fmt"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/db"
	"yama.io/yamaIterativeE/internal/util"
)

var (
	JavaSpringConfig         db.Config
	SpringMysqlConfig 		 db.Config
	SpringConsulConfig 		 db.Config
	SpringZipkinConfig 		 db.Config
	SpringActuatorConfig 	 db.Config
	SpringGRPCConfig  		 db.Config
	PythonConfig 			 db.Config
	GolangConfig 			 db.Config
)

type JavaSpringDynamicConfig struct {
	CONSUL_HOST  string
	CONSUL_PORT  string
	ZIPKIN_URL   string
	DATABASE_URL string
	DATABASE_UN  string
	DATABASE_PD  string
	HOST_NAME    string
	HOST_TAGS    string
	INSTANCE_ID  string
	APP_NAME     string
}

var JAVA_SPRING_DYNAMIC_CONFIG = JavaSpringDynamicConfig{
	CONSUL_HOST:  "spring.cloud.consul.host",
	CONSUL_PORT:  "spring.cloud.consul.port",
	ZIPKIN_URL:   "spring.zipkin.base-url",
	HOST_NAME:    "spring.cloud.consul.discovery.hostname",
	HOST_TAGS:    "spring.cloud.consul.discovery.tags",
	INSTANCE_ID:  "spring.cloud.consul.discovery.instance-id",
	APP_NAME:     "spring.application.name",
	DATABASE_URL: "spring.datasource.url",
	DATABASE_UN:  "spring.datasource.username",
	DATABASE_PD:  "spring.datasource.password",
}

func InitConfig() {
	initJavaSpringConfig()
	initSpringMysqlConfig()
	initSpringConsulConfig()
	initSpringGRPCConfig()
	initSpringActuatorConfig()
	initSpringZipkinConfig()
}

func initJavaSpringConfig() {
	JavaSpringConfig = db.GetJavaSpringConfig()
}
func initSpringMysqlConfig() {
	SpringMysqlConfig = db.GetSpringMysqlConfig()
	// mysql use global mysql
	(&SpringMysqlConfig).SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.DATABASE_URL, "jdbc:%s://%s:3306/%s")
}
func initSpringConsulConfig() {
	SpringConsulConfig = db.GetSpringConsulConfig()
	// consul use yamaiterativee proxy
	(&SpringConsulConfig).SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.CONSUL_HOST, util.GetLocalIPv4Address())
	(&SpringConsulConfig).SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.CONSUL_PORT, 4000)
}
func initSpringZipkinConfig() {
	SpringZipkinConfig = db.GetSpringZipkinConfig()
	// zipkin use global zipkin
	(&SpringZipkinConfig).SetConfigItem(JAVA_SPRING_DYNAMIC_CONFIG.ZIPKIN_URL, "http://%s:9411")
}
func initSpringGRPCConfig() {
	SpringGRPCConfig = db.GetSpringGRPCConfig()
}
func initSpringActuatorConfig() {
	SpringActuatorConfig = db.GetSpringActuatorConfig()
}

func GetApplicationConfig(c *context.Context) ([]byte, error) {
	key := c.Params(":key")
	var config *db.Config
	switch key {
	case "JAVA_SPRING":
		config = &JavaSpringConfig
		break
	case "SPRING_CONSUL":
		config = &SpringConsulConfig
		break
	case "SPRING_ACTUATOR":
		config = &SpringActuatorConfig
		break
	case "SPRING_MYSQL":
		config = &SpringMysqlConfig
		break
	case "SPRING_ZIPKIN":
		config = &SpringZipkinConfig
		break
	case "SPRING_GRPC":
		config = &SpringGRPCConfig
		break
	}
	if config == nil {
		return nil, fmt.Errorf("unsupport config key: %s", key)
	}

	var configItems []db.ConfigItem
	for _, v := range config.ConfigItems {
		if v.Displayable {
			configItems = append(configItems, v)
		}
	}
	data, err := json.Marshal(configItems)
	return data, err
}

func ResetApplicationNetwork() {

}