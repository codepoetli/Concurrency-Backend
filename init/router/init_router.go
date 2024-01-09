package router

import (
	"Concurrency-Backend/internal/controller"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// 用户注册、用户登录、用户上传视频、用户点赞、用户评论、用户follow其他用户

// InitRouterHertz 初始化hertz服务器路由
func InitRouterHertz(hertz *server.Hertz) {
	hertz.POST("/douyin/user/register/", controller.Register)
	
}
