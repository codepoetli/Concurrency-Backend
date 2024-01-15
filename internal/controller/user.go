package controller

import (
	"Concurrency-Backend/api"
	"Concurrency-Backend/internal/service"
	"Concurrency-Backend/utils/constants"
	"Concurrency-Backend/utils/jwt"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register 处理用户登录请求的RPC远程调用
func Register(context context.Context, requestContext *app.RequestContext) {
	var err error
	var user jwt.UserStruct
	if err = requestContext.BindAndValidate(&user); err != nil { // 格式错误
		requestContext.JSON(consts.StatusOK, api.UserLoginResponse{
			Response: api.Response{
				StatusCode: int32(api.InputFormatCheckErr),
				StatusMsg:  api.ErrorCodeToMsg[api.InputFormatCheckErr],
			},
		})
	}

	// 与service层通信
	err = service.GetUserServiceInstance().UserRegisterInfo(user.Username, user.Password)

	if err != nil {
		if errors.Is(errors.New(api.ErrorCodeToMsg[api.UserAlreadyExistErr]), err) {
			requestContext.JSON(consts.StatusOK, api.UserLoginResponse{
				Response: api.Response{
					StatusCode: int32(api.UserAlreadyExistErr),
					StatusMsg:  api.ErrorCodeToMsg[api.UserAlreadyExistErr],
				},
			})
		}
	}

	jwt.JwtMiddleware.LoginHandler(context, requestContext)
}

func UserInfo(c context.Context, ctx *app.RequestContext) {
	// 注意这里 test文件和接口文档有出入
	var err error
	var userId int64
	userId, err = jwt.GetUserId(c, ctx)
	if err != nil {
		ctx.JSON(consts.StatusOK, api.UserResponse{
			Response: api.Response{
				StatusCode: int32(api.TokenInvalidErr),
				StatusMsg:  api.ErrorCodeToMsg[api.TokenInvalidErr],
			},
		})
		return
	}

	// without JWT
	//userIdStr := ctx.Query("user_id")
	//userId, err := strconv.ParseInt(userIdStr, 10, 64)
	//if err != nil {
	//	ctx.JSON(consts.StatusOK, api.UserResponse{
	//		Response: api.Response{
	//			StatusCode: int32(api.InputFormatCheckErr),
	//			StatusMsg:  api.ErrorCodeToMsg[api.InputFormatCheckErr],
	//		},
	//	})
	//	return
	//}

	queryUser, err := service.GetUserServiceInstance().GetUserByUserId(userId)
	if errors.Is(constants.UserNotExistErr, err) {
		ctx.JSON(consts.StatusOK, api.UserResponse{
			Response: api.Response{
				StatusCode: int32(api.UserNotExistErr),
				StatusMsg:  api.ErrorCodeToMsg[api.UserNotExistErr],
			},
		})
		return
	}

	ctx.JSON(consts.StatusOK, api.UserResponse{
		Response: api.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		User: api.User{
			Id:            queryUser.UserID,
			Name:          queryUser.UserName,
			FollowCount:   queryUser.FollowCount,
			FollowerCount: queryUser.FollowerCount,
			IsFollow:      false,
		},
	})
}
