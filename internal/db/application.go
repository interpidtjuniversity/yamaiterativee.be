package db

import (
	"xorm.io/builder"
)

const (
	JAVA_SPRING string = "JAVA_SPRING"
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

func GetApplicationNetworkByOwnerAndRepo(owner, app string) (string, string){
	application := new(Application)
	exist, err := x.Table("application").Cols("network_name","network_ip").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Get(application)
	if err!=nil || !exist {
		return "", ""
	}
	return application.NetWorkName, application.NetWorkIP
}

func ApplicationIsExist(owner, app string) bool {
	exist, _ := x.Table("application").Cols("id").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Get(new(Application))
	return exist
}

func GetApplicationTypeByOwnerAndRepo(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("app_image").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.ApplicationImage
}

func GetApplicationRepoByOwnerAndRepo(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("repo_url").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.RepoUrl
}

func GetApplicationDevConfig(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("dev_config").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.DevConfig
}

func GetApplicationStableConfig(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("stable_config").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.StableConfig
}

func GetApplicationTestConfig(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("test_config").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.TestConfig
}

func GetApplicationPreConfig(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("pre_config").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.PreConfig
}

func GetApplicationProdConfig(owner, app string) string {
	application := new(Application)
	exist, err := x.Table("application").Cols("prod_config").Where(builder.Eq{"app_name": app}.And(builder.Eq{"owner":owner})).Limit(1).Get(application)
	if err!=nil || !exist {
		return ""
	}
	return application.ProdConfig
}