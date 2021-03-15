package db

import (
	"fmt"
	"testing"
)

func Test_UserInsert(t *testing.T) {
	NewEngine()

	user := User{Name: "tj-1752486-cy",LowerName: "tj-1752486-cy", IsActive: 1, Email: "120571672@qq.com", Passwd: "cy19991116"}
	if _, err := x.Insert(user); err!=nil{
		fmt.Print(err)
	}
}

func Test_UserQuery(t *testing.T) {
	NewEngine()

	//user := &User{ID: 0}
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