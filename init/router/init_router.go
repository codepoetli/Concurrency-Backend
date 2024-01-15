package router

import (
	"Concurrency-Backend/internal/controller"
	"Concurrency-Backend/utils/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// 用户注册、用户登录、用户上传视频、用户点赞、用户评论、用户follow其他用户

// InitRouterHertz 初始化hertz服务器路由
func InitRouterHertz(hertz *server.Hertz) {
	// public directory is used to serve static resources
	hertz.Static("/static", "./public")

	hertz.POST("/douyin/user/register/", controller.Register)
	hertz.POST("/douyin/user/login/", jwt.JwtMiddleware.LoginHandler)
	hertz.GET("/douyin/feed/", controller.Feed)

	// 鉴权authorization 中间件聚合的路由组上
	auth := hertz.Group("/douyin", jwt.JwtMiddleware.MiddlewareFunc())

	// basic apis
	auth.GET("/user/", controller.UserInfo)
	auth.POST("/publish/action/", controller.Publish)
	auth.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	auth.POST("/favorite/action/", controller.FavoriteAction)
	auth.GET("/favorite/list/", controller.FavoriteList)
	auth.POST("/comment/action/", controller.CommentAction)
	auth.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	//auth.POST("/relation/action/", controller.RelationAction)
	//auth.GET("/relation/follow/list/", controller.FollowList)
	//auth.GET("/relation/follower/list/", controller.FollowerList)
	//auth.GET("/relation/friend/list/", controller.FriendList)
	//auth.GET("/message/chat/", controller.MessageChat)
	//auth.POST("/message/action/", controller.MessageAction)

}
