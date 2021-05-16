package db

import "xorm.io/builder"

type User struct {
	Id        int64  `xorm:"id autoincr pk"`
	Name      string `xorm:"name"`
	LowerName string `xorm:"lower_name"`
	Email     string `xorm:"email"`
	Passwd    string `xorm:"passwd"`
	IsActive  int    `xorm:"is_active"`
	Avatar    string `xorm:"avatar"`
	ExtInfo   string `xorm:"ext_info"`
}

func GetAllUser() ([]string, error) {
	var users []User
	var names []string
	err := x.Cols("name").Where("id >= ?", 0).Find(&users)
	if err != nil {
		return nil, err
	}
	for _, user := range users{
		names = append(names, user.Name)
	}
	return names, nil

}

func BranchQueryUserByName(names []string) ([]*User, error){
	var users []*User
	err := x.Table("user").Where(builder.In("name", names)).Find(&users)
	return users, err
}

func GetUserAvatarByUserName(name string) (string, error) {
	user := new(User)
	exist, err := x.Table("user").Cols("avatar").Where(builder.Eq{"name":name}).Limit(1).Get(user)
	if !exist && err != nil {
		return "", err
	}
	return user.Avatar, nil
}