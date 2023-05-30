package config

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var Logger *logrus.Logger

func InitLog() {
	Logger = logrus.New()

	// 滚动切割日志
	fileLogger := &lumberjack.Logger{
		Filename:   "./example/logs/foo.log",
		MaxSize:    500,  // 单个日志文件大小(MB)
		MaxBackups: 3,    // 旧日志文件的数量
		MaxAge:     28,   // 日志存活时长(天)
		Compress:   true, // 是否对旧日志文件进行压缩
	}

	//设置日志输出(控制台输出和文件输出)
	Logger.SetOutput(io.MultiWriter(os.Stdout, fileLogger))

	// 设置日志级别为 Debug 级别
	Logger.SetLevel(logrus.DebugLevel)

	// 设置日志输出格式
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                      //键值对加引号
		TimestampFormat: "2006/01/02 15:04:05.000", //时间格式
		FullTimestamp:   true,
	})

}
