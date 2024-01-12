package constants

import (
	"Concurrency-Backend/api"
	"errors"
)

var (
	UserNotExistErr     = errors.New(api.ErrorCodeToMsg[api.UserNotExistErr])
	UserAlreadyExistErr = errors.New(api.ErrorCodeToMsg[api.UserAlreadyExistErr])
	InnerDataBaseErr    = errors.New(api.ErrorCodeToMsg[api.InnerDataBaseErr])
	CreateDataErr       = errors.New(api.ErrorCodeToMsg[api.CreateDataErr])
)
