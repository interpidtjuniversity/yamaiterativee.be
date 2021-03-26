package db

import (
	"fmt"
	"testing"
)

func Test_CreateIterationAction(t *testing.T) {
	NewEngine()

	action1 := IterationAction{
		ActorName: "interpidtjuniversity",
		State: "running",
		PipeLineId: 1,
		EnvGroup: 1,
		ActionInfo: "张启帆 给MR：#999999 的源分支提交代码触发了Pipeline #10000000 开发环境",
		AvatarSrc: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",
		ExtInfo: "this is extinfo",
	}

	action2 := IterationAction{
		ActorName: "interpidtjuniversity",
		State: "running",
		PipeLineId: 1,
		EnvGroup: 2,
		ActionInfo: "张启帆 给MR：#999999 的源分支提交代码触发了Pipeline #10000000 集成环境",
		AvatarSrc: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",
		ExtInfo: "this is extinfo",
	}

	action3 := IterationAction{
		ActorName: "interpidtjuniversity",
		State: "running",
		PipeLineId: 1,
		EnvGroup: 3,
		ActionInfo: "张启帆 给MR：#999999 的源分支提交代码触发了Pipeline #10000000 预发环境",
		AvatarSrc: "https://img.alicdn.com/tfs/TB1QS.4l4z1gK0jSZSgXXavwpXa-1024-1024.png",
		ExtInfo: "this is extinfo",
	}

	_,err := x.Insert(action1)
	_,err = x.Insert(action2)
	_,err = x.Insert(action3)
	if err!=nil {
		fmt.Print(err)
	}
}
