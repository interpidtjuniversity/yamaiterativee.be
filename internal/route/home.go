package route

import "yama.io/yamaIterativeE/internal/context"

func Home(ctx *context.Context) {
	ctx.Data["Name"] = "hello world"
	ctx.Data["PageIsAdmin"] = false
	ctx.Success("home")// 200 为响应码
}
