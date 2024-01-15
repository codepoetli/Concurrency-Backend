package controller

// 处理上传视频相关

import (
	"Concurrency-Backend/api"
	"Concurrency-Backend/internal/service"
	"Concurrency-Backend/utils/constants"
	"Concurrency-Backend/utils/jwt"
	"Concurrency-Backend/utils/logger"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"strconv"
	"time"
)

// Publish check token then save upload file to public directory
func Publish(c context.Context, ctx *app.RequestContext) {
	userId, err := jwt.GetUserId(c, ctx)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v,can't get user From token", time.Now())
		if errors.Is(constants.InvalidTokenErr, err) {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.TokenInvalidErr),
				StatusMsg:  api.ErrorCodeToMsg[api.TokenInvalidErr],
			})
		} else {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.InnerDataBaseErr),
				StatusMsg:  api.ErrorCodeToMsg[api.InnerDataBaseErr],
			})
		}
		return
	}
	logger.GlobalLogger.Printf("Time = %v,get User From loginUser = %v", time.Now(), userId)

	data, err := ctx.FormFile("data")
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v,can't get Video Data from post", time.Now())
		ctx.JSON(consts.StatusOK, api.Response{
			StatusCode: int32(api.GetDataErr),
			StatusMsg:  api.ErrorCodeToMsg[api.GetDataErr],
		})
		return
	}

	title := ctx.Query("title")
	err = service.GetPublishServiceInstance().PublishInfo(data, userId, title)
	if err != nil {
		ctx.JSON(consts.StatusOK, api.Response{
			StatusCode: int32(api.UploadFailErr),
			StatusMsg:  api.ErrorCodeToMsg[api.UploadFailErr],
		})
		return
	}
	ctx.JSON(consts.StatusOK, api.Response{
		StatusCode: 0,
	})
}

// PublishList 用户上传的视频列表
func PublishList(c context.Context, ctx *app.RequestContext) {
	loginUserId, err := jwt.GetUserId(c, ctx)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v,can't get user From token", time.Now())
		if errors.Is(constants.InvalidTokenErr, err) {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.TokenInvalidErr),
				StatusMsg:  api.ErrorCodeToMsg[api.TokenInvalidErr],
			})
		} else {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.InnerDataBaseErr),
				StatusMsg:  api.ErrorCodeToMsg[api.InnerDataBaseErr],
			})
		}
		return
	}

	userStr := ctx.Query("user_id")
	userId, err := strconv.ParseInt(userStr, 10, 64)
	logger.GlobalLogger.Printf("userId = %v", userId)
	if err != nil {
		ctx.JSON(consts.StatusOK, api.Response{
			StatusCode: int32(api.InputFormatCheckErr),
			StatusMsg:  api.ErrorCodeToMsg[api.InputFormatCheckErr],
		})
		return
	}

	videoList, err := service.GetPublishServiceInstance().PublishListInfo(userId, loginUserId)
	ctx.JSON(consts.StatusOK, api.VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
