package db

type Application struct {
	ID         int64      `xorm:"id autoincr pk"`
	AppName    string     `xorm:"app_name"`
	NetWork    string     `xorm:"net_work"`
	Owner      string     `xorm:"owner"`
	RepoUrl    string     `xorm:"repo_url"`
	TestConfig string     `xorm:"test_config"`
	ProdConfig string     `xorm:"prod_config"`
	Users      []int64    `xorm:"users"`
}

func NewApplication() {
	//1. check repository and new repository
	//2. allocate network
	//3. allocate HostPort range if needed(it will be fast for api)
	//4. write db
}


