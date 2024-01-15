package jwt

import (
	"Concurrency-Backend/internal/model"
	"Concurrency-Backend/internal/service"
	"Concurrency-Backend/utils/constants"
	"Concurrency-Backend/utils/logger"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

// GetUserId 从token中获取userid
func GetUserId(content context.Context, requestContext *app.RequestContext) (int64, error) {
	user, exists := requestContext.Get(IdentityKey)
	if !exists {
		return 0, constants.InvalidTokenErr
	}

	loginUserInfo := user.(*model.User)
	logger.GlobalLogger.Printf("Time = %v, In UserInfo, Got Login Username =%v", time.Now(), loginUserInfo.UserName)
	loginUserInfo, err := service.GetUserServiceInstance().GetUserIdByUserName(loginUserInfo.UserName)
	return loginUserInfo.UserID, err
}
