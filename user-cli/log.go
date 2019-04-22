package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)


var llog *logrus.Entry
var levelMap = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}


func init() {
	rlog := logrus.New()
	// 以JSON格式为输出，代替默认的ASCII格式
	rlog.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	rlog.SetOutput(os.Stdout)
	// 设置日志等级,先从环境变量中读取，然后判断下
	var logLevel string
	if os.Getenv("loglevel")!= ""{
		logLevel = strings.ToLower(os.Getenv("loglevel"))
	}

	if logLevel != "debug" && logLevel != "info" && logLevel != "warn" && logLevel != "error" && logLevel != "fatal" && logLevel != "panic" {
		logLevel = "debug"
	}
	rlog.SetLevel(levelMap[logLevel])
	// 每次打印都带上这个字段
	llog = rlog.WithField("SERVICE","[user-cli]")
}
