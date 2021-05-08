package db

import (
	"xorm.io/builder"
)

type AppType int
const (
	JAVA_SPRING AppType = iota
)

type Application struct {
	ID                        int64    `xorm:"id autoincr pk"`
	AppName                   string   `xorm:"app_name"`
	NetWorkIP                 string   `xorm:"network_ip"`
	NetWorkName               string   `xorm:"network_name"`
	Owner                     string   `xorm:"owner"`
	RepoUrl                   string   `xorm:"repo_url"`
	DevInstances              []string `xorm:"dev_instances"`
	StableInstances           []string `xorm:"stable_instances"`
	TestInstances             []string `xorm:"test_instances"`
	PreInstances              []string `xorm:"pre_instances"`
	ProdInstances             []string `xorm:"prod_instances"`
	TestConfig                string   `xorm:"test_config"`
	ProdConfig                string   `xorm:"prod_config"`
	DevConfig                 string   `xorm:"dev_config"`
	PreConfig                 string   `xorm:"pre_config"`
	StableConfig              string   `xorm:"stable_config"`
	Users                     []string `xorm:"users"`
	AvatarURL                 string   `xorm:"avatar"`
	ApplicationImage          string   `xorm:"app_image"`
	ApplicationRegistry       string   `xorm:"app_registry"`
	ApplicationTrace          string   `xorm:"app_trace"`
	ApplicationAuth           string   `xorm:"app_auth"`
	ApplicationDataBase       string   `xorm:"app_database"`
	ApplicationBusinessDomain string   `xorm:"app_businessDomain"`
	ApplicationDomainName     string   `xorm:"app_domain"`
	Description               string   `xorm:"description"`
}

func InsertApplication(application *Application) error{
	_, err := x.Table("application").Insert(application)
	return err
}

func GetApplicationByUser(owner string) ([]*Application, error) {
	var applications []*Application
	err := x.Table("application").Where(builder.Eq{"owner": owner}).Find(&applications)
	return applications, err
}

func GetApplicationByParticipant(user string) ([]*Application, error) {
	var applications []*Application
	var filterApplicationIds []int64
	err := x.Table("application").Cols("id", "users").Find(&applications)
	for _, app := range applications {
		for _, u := range app.Users {
			if u == user {
				filterApplicationIds = append(filterApplicationIds, app.ID)
				break
			}
		}
	}
	applications = applications[:0]
	err = x.Table("application").Where(builder.In("id", filterApplicationIds)).Find(&applications)
	return applications, err
}
