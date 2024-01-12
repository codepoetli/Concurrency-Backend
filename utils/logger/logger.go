package logger

import (
	initialization "Concurrency-Backend/init"
	"Concurrency-Backend/utils/files"
	"github.com/rs/zerolog"
	"os"
)

// GlobalLogger 全局logger
var GlobalLogger zerolog.Logger

// InitLogger 初始化GlobalLogger
func InitLogger(config initialization.LogConfig) {
	var err error
	var file *os.File
	if config.LogFileWritten { // 配置了log文件
		if exists, _ := files.PathExists(config.LogFilePath); exists {
			file, err = os.OpenFile(config.LogFilePath, os.O_APPEND, 0666)
		} else {
			file, err = os.Create(config.LogFilePath)
		}
		if err != nil {
			panic(err)
		}
		GlobalLogger = zerolog.New(file)
	} else {
		GlobalLogger = initialization.GetStdOutLogger()
	}
	GlobalLogger.Level(zerolog.InfoLevel)
}
