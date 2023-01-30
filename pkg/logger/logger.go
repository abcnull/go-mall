package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

func init() { // 每天获取的新路径使不同的，怎么保证 init 每天都会被执行
	// 拿到输出日志的新路径
	src, _ := setOutputFile()

	// 每次 logger 的 output 路径得更新一下
	if Logger != nil {
		Logger.Out = src
		return
	}

	// 首次就需要实例化
	Logger = logrus.New()              // 生成一个新 Logger
	Logger.Out = src                   // 设置输出路径
	Logger.SetLevel(logrus.DebugLevel) // 设置日志级别
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// 设置日志输出路径
func setOutputFile() (*os.File, error) {
	// 日志 logs 路径
	var logFilePath string                  // 日志路径
	if dir, err := os.Getwd(); err != nil { // 获取工作根目录
		logFilePath = dir + "/logs/"
	}

	_, err := os.Stat(logFilePath) // todo: 不太理解 os 操作
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil { // 创建失败这个 log 路径
			log.Println(err.Error())
			return nil, err
		}
	}

	// 加上日志名的路径
	logFileName := time.Now().Format("2006-01-02") + ".log" // 日志名称
	fileName := path.Join(logFilePath, logFileName)         // todo: path 包不太熟悉
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(fileName, 0777); err != nil { // 创建失败这个 log 路径
			log.Println(err.Error())
			return nil, err
		}
	}

	// 日志写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend) // todo: 这里不太熟悉
	if err != nil {
		return nil, err
	}

	return src, nil
}
