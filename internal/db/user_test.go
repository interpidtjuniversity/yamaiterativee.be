package db

import (
	"fmt"
	"testing"
	"xorm.io/builder"
	"yama.io/yamaIterativeE/internal/util"
)

func Test_UserInsert(t *testing.T) {
	NewEngine()

	user := User{Name: "tj-1752486-cy",LowerName: "tj-1752486-cy", IsActive: 1, Email: "120571672@qq.com", Passwd: "cy19991116"}
	_,_= x.Insert(&user)
	fmt.Print(user.Id)
	fmt.Print(user.Id)
}

func Test_UserQuery(t *testing.T) {
	NewEngine()

	//user := &User{Id: 0}
	//if _,err := x.Get(user); err!=nil {
	//	fmt.Print(err)
	//}
	//fmt.Print(user)

	user1 := &User{}
	has, _ := x.ID(int64(1)).Get(user1)
	if has{
		fmt.Print(user1)
	} else {
		fmt.Print("not exist")
	}
}

func Test_BranchQuery(t *testing.T) {
	NewEngine()

	var users []*User
    err := x.Where(builder.Eq{"email":"120571672@qq.com"}).Find(&users)

    var ids []int64
    stream, _ := util.New(users)
    err = stream.Map(func(user *User)int64 {
    	return user.Id
	}).ToSlice(&ids)

	if err!=nil {
		fmt.Print(err)
	}
	fmt.Print(ids)
}

func Test_InSet(t *testing.T) {
	NewEngine()

	var users []*User
	err := x.Where(builder.In("id", []int64{0,0,0,0,0})).Find(&users)
	if err!=nil{
		fmt.Print(err)
	}
	stream, _ := util.New(users)
	var names []string
	stream.Map(func(user *User)string {
		return user.Name
	}).ToSlice(&names)
	fmt.Print(names)

}