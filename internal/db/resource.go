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

func InsertResource(resource *Resource) error {
	r := new(Resource)
	exist, err := x.Table("resource").Where(builder.Eq{"name": resource.Name}).Get(r)
	if exist {
		r.Value = resource.Value
		_, err :=x.Table("resource").Where(builder.Eq{"name": resource.Name}).Update(r)
		return err
	}
	if err != nil {
		return err
	}
	_, err = x.Table("resource").Insert(resource)
	return err
}