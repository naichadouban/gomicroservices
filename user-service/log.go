package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var llog *logrus.Logger
var levelMap = map[string]logrus.Level{
	"info":  logrus.InfoLevel,
	"debug": logrus.DebugLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

func init() {
	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	logrus.SetOutput(os.Stdout)
	// 设置日志等级,先从环境变量中读取，然后判断下
	logLevel := strings.ToLower(os.Getenv("loglevel"))
	if logLevel != "debug" && logLevel != "info" && logLevel != "warn" && logLevel != "error" && logLevel != "fatal" && logLevel != "panic" {
		logLevel = "debug"
	}
	logrus.SetLevel(levelMap[logLevel])
}
