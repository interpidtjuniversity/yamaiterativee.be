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

// same as config table!!!
var JAVA_SPRING_CONFIG = Config{
	ApplicationType: "JAVA_SPRING",
	ConfigItems: []ConfigItem{
		// default config and can not be changed
		{"spring.sleuth.enabled", true, false, true},
		{"grpc.client.cloud-grpc-server-consul.enableKeepAlive", true, false, true},
		{"grpc.client.cloud-grpc-server-consul.keepAliveWithoutCalls", true, false, true},
		{"grpc.client.cloud-grpc-server-consul.negotiationType", "plaintext", false, true},
		{"management.endpoints.web.exposure.include", "*", false, true},
		{"management.server.port", 8088, false, true},
		{"management.endpoint.health.show-details", "always", false, true},
		{"management.endpoint.serviceregistry.enabled", true, false, true},
		{"server.port", 8080, false, true},
		{"spring.zipkin.enabled", true, false, true},
		{"spring.sleuth.sampler.probability", 1, false, true},
		{"spring.sleuth.grpc.enabled", true, false, true},
		{"spring.cloud.consul.discovery.register", true, false, true},
		{"spring.cloud.consul.discovery.port", 10000, false, true},
		{"spring.datasource.username", "root", false, true},
		{"spring.datasource.password", 123456, false, true},

		// dynamic config and change with yamaiterativee env
		{"spring.datasource.url", "-", false, true},
		{"spring.cloud.consul.host", "-", false, true},
		{"spring.cloud.consul.port", "-", false, true},
		{"spring.zipkin.base-url", "-", false, true},

		// custom config and can be changed only once
		{"spring.application.name", "-", true, true},

		// deploy config and it value will be write when in deploying
		{"spring.cloud.consul.discovery.hostname", "-", false, false},
		{"spring.cloud.consul.discovery.tags", "-", false, false},
		{"spring.cloud.consul.discovery.instance-id", "-", false, false},
	},
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