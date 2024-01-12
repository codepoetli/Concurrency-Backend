package router

import (
	"Concurrency-Backend/internal/controller"
	"Concurrency-Backend/utils/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// 用户注册、用户登录、用户上传视频、用户点赞、用户评论、用户follow其他用户

// InitRouterHertz 初始化hertz服务器路由
func InitRouterHertz(hertz *server.Hertz) {
	// public directory is used to serve static resources 微服务要删去
	hertz.Static("/static", "./public")

	hertz.POST("/douyin/user/register/", controller.Register)
	hertz.POST("/douyin/user/login/", jwt.JwtMiddleware.LoginHandler)
	// hertz.GET("/douyin/feed/", controller.Feed) todo

}
