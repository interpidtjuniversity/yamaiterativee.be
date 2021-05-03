package db

type Application struct {
	ID              int64    `xorm:"id autoincr pk"`
	AppName         string   `xorm:"app_name"`
	NetWork         string   `xorm:"network"`
	Owner           string   `xorm:"owner"`
	RepoUrl         string   `xorm:"repo_url"`
	DevInstances    []string `xorm:"dev_instances"`
	StableInstances []string `xorm:"stable_instances"`
	TestInstances   []string `xorm:"test_instances"`
	PreInstances    []string `xorm:"pre_instances"`
	ProdInstances   []string `xorm:"prod_instances"`
	TestConfig      string   `xorm:"test_config"`
	ProdConfig      string   `xorm:"prod_config"`
	DevConfig       string   `xorm:"dev_config"`
	Users           []string `xorm:"users"`
}

func NewApplication() {
	//1. check repository and new repository
	//2. allocate network
	//3. allocate HostPort range if needed(it will be fast for api)
	//4. write db
}


