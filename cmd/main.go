package main

import (
	"Concurrency-Backend/init/router"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func InitAll() {

}

func main() {
	hServer := server.Default() // 修改端口 todo

	router.InitRouterHertz(hServer)
	hServer.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	hServer.Spin() // 运行 Hertz 服务器，接收到退出信号后可退出服务
}
