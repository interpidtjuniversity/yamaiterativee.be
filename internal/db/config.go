package db

import (
	"xorm.io/builder"
)

type Config struct {
	ID                  int64        `xorm:"id autoincr pk"`
	ApplicationType     string       `xorm:"app_type"`
	ConfigItems         []ConfigItem `xorm:"config_items"`
}
func (c *Config) SetConfigItem(key string, value interface{}) {
	for i:=0; i<len(c.ConfigItems); i++ {
		if c.ConfigItems[i].Key == key {
			c.ConfigItems[i].Value = value
			return
		}
	}
	newItem := ConfigItem{Key: key, Value: value}
	c.ConfigItems = append(c.ConfigItems, newItem)
}

func (c *Config) GetConfigItem(key string) interface{} {
	for i:=0; i<len(c.ConfigItems); i++ {
		if c.ConfigItems[i].Key == key {
			return c.ConfigItems[i].Value
		}
	}
	return nil
}

type ConfigItem struct {
	Key         string      `json:"key"`
	Value       interface{} `json:"value"`
	Changeable  bool        `json:"changeable"`
	Displayable bool        `json:"displayable"`
}

var SPRING_MYSQL_CONFIG = Config{
	ApplicationType: "SPRING_MYSQL",
	ConfigItems: []ConfigItem{
		{"spring.datasource.username", "root", false, true},
		{"spring.datasource.password", 123456, false, true},
		{"spring.datasource.url", "-", false, true},
	},
}

var SPRING_CONSUL_CONFGI = Config{
	ApplicationType: "SPRING_CONSUL",
	ConfigItems: []ConfigItem{
		{"spring.cloud.consul.host", "-", false, true},
		{"spring.cloud.consul.port", "-", false, true},
		{"spring.cloud.consul.discovery.hostname", "-", false, false},
		{"spring.cloud.consul.discovery.tags", "-", false, false},
		{"spring.cloud.consul.discovery.instance-id", "-", false, false},
		{"spring.cloud.consul.discovery.register", true, false, true},
		{"spring.cloud.consul.discovery.port", 10000, false, true},
	},
}

var SPRING_GRPC_CONFIG = Config{
	ApplicationType: "SPRING_GRPC",
	ConfigItems: []ConfigItem{
		{"grpc.client.cloud-grpc-server-consul.enableKeepAlive", true, false, true},
		{"grpc.client.cloud-grpc-server-consul.keepAliveWithoutCalls", true, false, true},
		{"grpc.client.cloud-grpc-server-consul.negotiationType", "plaintext", false, true},
	},
}

var SPRING_ACTUATOR_CNFIG = Config{
	ApplicationType: "SPRING_ACTUATOR",
	ConfigItems: []ConfigItem{
		{"management.endpoints.web.exposure.include", "*", false, true},
		{"management.server.port", 8088, false, true},
		{"management.endpoint.health.show-details", "always", false, true},
		{"management.endpoint.serviceregistry.enabled", true, false, true},
	},
}

var SPRING_ZIPKIN_CONFIG = Config{
	ApplicationType: "SPRING_ZIPKIN",
	ConfigItems: []ConfigItem{
		{"spring.zipkin.enabled", true, false, true},
		{"spring.sleuth.sampler.probability", 1, false, true},
		{"spring.sleuth.grpc.enabled", true, false, true},
		{"spring.sleuth.enabled", true, false, true},
		{"spring.zipkin.base-url", "-", false, true},
	},
}

// same as config table!!!
var JAVA_SPRING_CONFIG = Config{
	ApplicationType: "JAVA_SPRING",
	ConfigItems: []ConfigItem{
		{"server.port", 8080, false, true},
		{"spring.application.name", "-", true, true},
	},
}

func InsertSpringMysqlConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_MYSQL"}).Get(config)
	if exist {
		SPRING_MYSQL_CONFIG.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_MYSQL"}).Update(&SPRING_MYSQL_CONFIG)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&SPRING_MYSQL_CONFIG)
	return err
}

func InsertSpringConsulConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_CONSUL"}).Get(config)
	if exist {
		SPRING_CONSUL_CONFGI.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_CONSUL"}).Update(&SPRING_CONSUL_CONFGI)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&SPRING_CONSUL_CONFGI)
	return err
}

func InsertSpringGRPCConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_GRPC"}).Get(config)
	if exist {
		SPRING_GRPC_CONFIG.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_GRPC"}).Update(&SPRING_GRPC_CONFIG)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&SPRING_GRPC_CONFIG)
	return err
}

func InsertSpringActuatorConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_ACTUATOR"}).Get(config)
	if exist {
		SPRING_ACTUATOR_CNFIG.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_ACTUATOR"}).Update(&SPRING_ACTUATOR_CNFIG)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&SPRING_ACTUATOR_CNFIG)
	return err
}

func InsertSpringZipkinConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_ZIPKIN"}).Get(config)
	if exist {
		SPRING_ZIPKIN_CONFIG.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "SPRING_ZIPKIN"}).Update(&SPRING_ZIPKIN_CONFIG)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&SPRING_ZIPKIN_CONFIG)
	return err
}

func InsertJavaSpringConfig() error {
	config := new(Config)
	exist, err := x.Table("config").Where(builder.Eq{"app_type": "JAVA_SPRING"}).Get(config)
	if exist {
		JAVA_SPRING_CONFIG.ID = config.ID
		_, err := x.Table("config").Where(builder.Eq{"app_type": "JAVA_SPRING"}).Update(&JAVA_SPRING_CONFIG)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("config").Insert(&JAVA_SPRING_CONFIG)
	return err
}

func GetJavaSpringConfig() Config {
	if JAVA_SPRING_CONFIG.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "JAVA_SPRING"}).Get(config)
		JAVA_SPRING_CONFIG.ID = config.ID
		return *config
	}
	return JAVA_SPRING_CONFIG
}

func GetSpringMysqlConfig() Config {
	if SPRING_MYSQL_CONFIG.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "SPRING_MYSQL"}).Get(config)
		SPRING_MYSQL_CONFIG.ID = config.ID
		return *config
	}
	return SPRING_MYSQL_CONFIG
}

func GetSpringConsulConfig() Config {
	if SPRING_CONSUL_CONFGI.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "SPRING_CONSUL"}).Get(config)
		SPRING_CONSUL_CONFGI.ID = config.ID
		return *config
	}
	return SPRING_CONSUL_CONFGI
}

func GetSpringActuatorConfig() Config {
	if SPRING_ACTUATOR_CNFIG.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "SPRING_ACTUATOR"}).Get(config)
		SPRING_ACTUATOR_CNFIG.ID = config.ID
		return *config
	}
	return SPRING_ACTUATOR_CNFIG
}

func GetSpringGRPCConfig() Config {
	if SPRING_GRPC_CONFIG.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "SPRING_GRPC"}).Get(config)
		SPRING_GRPC_CONFIG.ID = config.ID
		return *config
	}
	return SPRING_GRPC_CONFIG
}

func GetSpringZipkinConfig() Config {
	if SPRING_ZIPKIN_CONFIG.ID == 0 {
		config := new(Config)
		_, _ = x.Table("config").Where(builder.Eq{"app_type": "SPRING_ZIPKIN"}).Get(config)
		SPRING_ZIPKIN_CONFIG.ID = config.ID
		return *config
	}
	return SPRING_ZIPKIN_CONFIG
}