package api

// ErrorType ：不同类型error对于不同的code和message
type ErrorType uint64

const (
	UploadFailErr     ErrorType = 10001
	SavingFailErr     ErrorType = 10002
	VideoFormationErr ErrorType = 10003
	VideoSizeErr      ErrorType = 10004
	NoVideoErr        ErrorType = 10005

	InnerDataBaseErr    ErrorType = 10101
	CreateDataErr       ErrorType = 10106
	TokenInvalidErr     ErrorType = 10107
	UserNotExistErr     ErrorType = 10108
	UserAlreadyExistErr ErrorType = 10109
	RecordNotExistErr   ErrorType = 10111
	RecordNotMatchErr   ErrorType = 10113

	UnKnownActionType   ErrorType = 10202
	InputFormatCheckErr ErrorType = 10203
	GetDataErr          ErrorType = 10204
)

var ErrorCodeToMsg = map[ErrorType]string{
	UploadFailErr:     "Fail to upload File",
	SavingFailErr:     "Fail to save file",
	VideoFormationErr: "Video formation error",
	VideoSizeErr:      "Video size larger than expected",
	NoVideoErr:        "No video matches the requirement",

	InnerDataBaseErr:    "Inner database error",
	CreateDataErr:       "Create data error",
	TokenInvalidErr:     "Invalid Token",
	UserNotExistErr:     "User not exist",
	UserAlreadyExistErr: "用户名已存在",
	RecordNotExistErr:   "数据不存在",
	RecordNotMatchErr:   "Record doesn't match",

	UnKnownActionType:   "Unknown Action Type",
	InputFormatCheckErr: "Input formation error",
	GetDataErr:          "Fail to get data from context",
}
