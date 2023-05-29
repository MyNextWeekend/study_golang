package config

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLog() {
	Logger = logrus.New()
	// 创建日志文件
	//file, err := os.OpenFile("./example/logs/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	panic(err.Error())
	//}
	//defer file.Close()

	// 设置日志输出 (控制台输出和文件输出)
	//Logger.SetOutput(io.MultiWriter(os.Stdout, file))

	// 设置日志级别为 Debug 级别
	Logger.SetLevel(logrus.DebugLevel)

	// 设置日志输出格式
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  //键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp:   true,
	})

}
