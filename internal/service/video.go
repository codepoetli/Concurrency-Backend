package service

import (
	"Concurrency-Backend/api"
	"Concurrency-Backend/internal/dao"
	"Concurrency-Backend/internal/model"
	"Concurrency-Backend/utils/constants"
)

// 构造Video切片, userId是当前登录的userId
func newVideoList(userId int64, videos []*model.Video) ([]api.Video, error) {
	videoList := make([]api.Video, len(videos))
	for i, v := range videos {
		userInfo, err := GetUserServiceInstance().GetUserByUserId(v.UserID)
		isFavor, err := dao.GetFavoriteDaoInstance().CheckFavorite(userId, v.VideoID) // 获取点赞数据
		if err != nil {
			return nil, constants.InnerDataBaseErr
		}
		videoList[i] = api.Video{
			Id: v.VideoID,
			Author: api.User{
				Id:            userInfo.UserID,
				Name:          userInfo.UserName,
				FollowCount:   userInfo.FollowCount,
				FollowerCount: userInfo.FollowerCount,
				IsFollow:      false,
			},
			PlayUrl:       v.PlayURL,
			CoverUrl:      v.CoverURL,
			FavoriteCount: int64(v.FavoriteCount),
			CommentCount:  int64(v.CommentCount),
			IsFavorite:    isFavor,
		}
	}
	return videoList, nil
}
