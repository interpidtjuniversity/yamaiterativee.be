package db

import "xorm.io/builder"

/**
k,v storage in database
*/
type Resource struct {
	Id    int64  `xorm:"id autoincr pk"`
	Name  string `xorm:"name"`  // server name GLOBAL_MYSQL GLOBAL_CONSUL GLOBAL_ZIPKIN
	Value string `xorm:"value"` // server ip
}

func GetResourceByName(name string) (*Resource, error) {
	resource := new(Resource)
	exist, err := x.Table("resource").Where(builder.Eq{"name": name}).Get(resource)
	if err!=nil || !exist {
		return nil, err
	}
	return resource, nil
}