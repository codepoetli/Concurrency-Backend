package main

import (
	initialization "Concurrency-Backend/init"
	"Concurrency-Backend/init/router"
	"Concurrency-Backend/utils/logger"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 用于单机的极简版抖音后端程序

// initAll 初始化所有的部分
func initAll() {
	initialization.InitConfig()
	// TODO 这些注释掉的部分之后再慢慢加上
	// initialization.InitDB()
	// initialization.InitOSS()
	// initialization.InitRDB()
	logger.InitLogger(initialization.LogConf)

	// jwt.InitJwt()
}

func main() {
	hServer := server.Default() // 修改端口 todo

	router.InitRouterHertz(hServer)
	hServer.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	hServer.Spin() // 运行 Hertz 服务器，接收到退出信号后可退出服务
}
