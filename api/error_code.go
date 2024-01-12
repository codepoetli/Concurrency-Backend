package api

// ErrorType ：不同类型error对于不同的code和message
type ErrorType uint64

const (
	InnerDataBaseErr    ErrorType = 10101
	CreateDataErr       ErrorType = 10106
	UserNotExistErr     ErrorType = 10108
	UserAlreadyExistErr ErrorType = 10109
	InputFormatCheckErr ErrorType = 10203
)

var ErrorCodeToMsg = map[ErrorType]string{
	InnerDataBaseErr:    "Inner database error",
	CreateDataErr:       "Create data error",
	UserNotExistErr:     "User not exist",
	UserAlreadyExistErr: "用户名已存在",

	InputFormatCheckErr: "Input formation error",
}
