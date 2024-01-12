package controller

import (
	"Concurrency-Backend/api"
	"Concurrency-Backend/internal/service"
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
