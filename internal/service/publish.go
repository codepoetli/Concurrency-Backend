package service

import (
	"Concurrency-Backend/api"
	initialization "Concurrency-Backend/init"
	"Concurrency-Backend/internal/dao"
	"Concurrency-Backend/internal/model"
	"Concurrency-Backend/internal/oss"
	"Concurrency-Backend/utils/constants"
	"Concurrency-Backend/utils/files"
	"Concurrency-Backend/utils/idGenerator"
	"Concurrency-Backend/utils/logger"
	"mime/multipart"
	"path"
	"strconv"
	"sync"
	"time"
)

// publishService 与publish相关的操作集合
type publishService struct{}

var (
	publishServiceInstance *publishService
	publishOnce            sync.Once
)

func GetPublishServiceInstance() *publishService {
	publishOnce.Do(func() {
		publishServiceInstance = &publishService{}
	})
	return publishServiceInstance
}

func getUploadPath(userId int64, fileName string) string {
	return initialization.OssConf.BucketDirectory + "/" + strconv.FormatInt(userId, 10) + "/" + fileName
}

// getUploadURL 得到一名用户对应的云端存储路径 oss
func getUploadURL(userId int64, fileName string) string {
	return "https://" + initialization.OssConf.Bucket + "." + initialization.OssConf.Url + "/" + getUploadPath(userId, fileName)
}

// 上传视频到oss
func (p *publishService) uploadVideoToOSS(data *multipart.FileHeader, userId int64, filename string) error {
	// src, err := data.Open()
	// if err != nil {
	// 	logger.GlobalLogger.Printf("Error in OpenData: %v", err.Error())
	// 	return err
	// }
	// defer src.Close()

	// // 先将文件流上传至BucketDirectory目录下
	// err = oss.UploadFromReader(getUploadPath(userId, filename), src)
	// if err != nil {
	// 	logger.GlobalLogger.Printf("Error in UploadFromReader: %v", err.Error())
	// 	return err
	// }

	return nil
}

// 上传封面图
func (p *publishService) uploadCoverToOSS(userId int64, filepath, filename string) error {
	if err := oss.UploadFromFile(getUploadPath(userId, filename), filepath); err != nil {
		logger.GlobalLogger.Printf("Error in UploadFromFile: %v", err.Error())
		return err
	}

	return nil
}

// PublishInfo service层上传user的一个视频
func (p *publishService) PublishInfo(data *multipart.FileHeader, userId int64, title string) error {
	logger.GlobalLogger.Printf("title = %v", title)
	fileName := data.Filename
	logger.GlobalLogger.Printf("fileName = %v", fileName)
	//首先检查video的扩展名与大小
	if !files.CheckFileExt(fileName) {
		return constants.VideoFormatErr
	}
	if !files.CheckFileSize(data.Size) {
		return constants.VideoSizeErr
	}

	logger.GlobalLogger.Print("Start Saving")

	//然后将文件保存至本地
	saveDir := path.Join(initialization.VideoConf.SavePath, strconv.FormatInt(userId, 10))
	videoName, err := files.SaveFileToLocal(saveDir, data)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v, Saving Video Error = %v", time.Now(), err.Error())
		return constants.SavingFailErr
	}

	coverName := ""
	// 截取视频的第一帧作为cover
	// saveVideo := saveDir + "/" + videoName
	// coverName := files.GetFileNameWithoutExt(videoName) + "_cover" + ".jpeg"
	// saveCover := saveDir + "/" + coverName
	// err = files.ExtractCoverFromVideo(saveVideo, saveCover)
	// if err != nil {
	// 	logger.GlobalLogger.Printf("Time = %v, Extracting Cover Error = %v", time.Now(), err.Error())
	// 	return constants.SavingFailErr
	// }
	// coverName := files.GetFileNameWithoutExt(videoName) + "_cover" + ".jpeg"
	// saveCover := saveDir + "/" + coverName

	//上传视频与封面 怎么回滚？
	logger.GlobalLogger.Print("Saving Complete, Start Uploading")
	err = p.uploadVideoToOSS(data, userId, videoName)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v, Upload Video Error = %v", time.Now(), err.Error())
		return constants.UploadFailErr
	}
	// err = p.uploadCoverToOSS(userId, saveCover, coverName)
	// if err != nil {
	// 	logger.GlobalLogger.Printf("Time = %v, Upload Cover Error = %v", time.Now(), err.Error())
	// 	return constants.UploadFailErr
	// }

	//写入数据库
	video := &model.Video{
		VideoID:       idGenerator.GenerateVideoId(),
		VideoName:     title,
		UserID:        userId,
		FavoriteCount: 0,
		CommentCount:  0,
		PlayURL:       getUploadURL(userId, videoName),
		CoverURL:      getUploadURL(userId, coverName),
	}
	err = dao.GetVideoDaoInstance().CreateVideo(video)
	return err
}

// PublishListInfo service层获得用户userId所有发表过的视频
func (p *publishService) PublishListInfo(userId, loginUserId int64) ([]api.Video, error) {
	var err error
	videoList, err := dao.GetVideoDaoInstance().GetPublishList(userId)
	if err != nil {
		return nil, err
	}
	apiVideos, err := newVideoList(loginUserId, videoList)
	return apiVideos, nil
}
