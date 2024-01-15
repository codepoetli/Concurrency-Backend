package service

import (
	initialization "Concurrency-Backend/init"
	"Concurrency-Backend/internal/dao"
	"Concurrency-Backend/internal/model"
	"Concurrency-Backend/utils/constants"
	"Concurrency-Backend/utils/idGenerator"
	"Concurrency-Backend/utils/logger"
	"Concurrency-Backend/utils/md5"
	"errors"
	"github.com/rs/zerolog/log"
	"sync"
)

// UserService 与用户相关的操作使用的结构体 单例模式的简单实现
type UserService struct{}

var (
	userServiceInstance *UserService
	userOnce            sync.Once
)

func GetUserServiceInstance() *UserService {
	userOnce.Do(func() {
		userServiceInstance = &UserService{}
	})
	return userServiceInstance
}

func (u *UserService) UserRegisterInfo(username, password string) error {
	var err error
	userInfo, err := dao.GetUserDaoInstance().GetUserByUsername(username)

	if errors.Is(constants.InnerDataBaseErr, err) {
		logger.GlobalLogger.Error().Caller().Str("用户注册失败", err.Error())
		return err
	}

	if userInfo != nil {
		logger.GlobalLogger.Error().Caller().Str("用户名已存在", err.Error())
		return constants.UserAlreadyExistErr
	}

	userId := idGenerator.GenerateUserId()
	logger.GlobalLogger.Info().Int64("userId ", userId)

	user := &model.User{
		UserID:   userId,
		UserName: username,
	}

	if initialization.UserConf.PasswordEncrpted {
		user.PassWord = md5.MD5(password)
	} else {
		user.PassWord = password
	}

	err = dao.GetUserDaoInstance().CreateUser(user) // 调用dao层写入

	if err != nil {
		log.Error().Caller().Str("用户注册错误", err.Error())
		return constants.CreateDataErr
	}
	return nil
}

// GetUserByUserId 通过userid得到user
func (u *UserService) GetUserByUserId(userId int64) (*model.User, error) {
	userInfo, err := dao.GetUserDaoInstance().GetUserByUserId(userId)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v, 寻找数据失败, err = %s", err.Error())
		return nil, err
	}
	return userInfo, err
}

// GetUserIdByUserName 通过username得到user
func (u *UserService) GetUserIdByUserName(username string) (*model.User, error) {
	userInfo, err := dao.GetUserDaoInstance().GetUserByUsername(username)
	if err != nil {
		logger.GlobalLogger.Printf("Time = %v, 寻找数据失败, err = %s", err.Error())
	}
	return userInfo, err
}
