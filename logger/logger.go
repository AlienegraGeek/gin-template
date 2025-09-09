package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

var Log *logrus.Logger

func Init() {
	Log = logrus.New()

	// 设置输出目标（默认 stdout）
	Log.Out = os.Stdout

	// 显示调用文件和行号
	Log.SetReportCaller(true)

	// 设置日志格式：JSON
	//logrus.SetFormatter(&logrus.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//	PrettyPrint:     true, // 格式化输出
	//})
	// 设置日志格式：Text
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true, // 强制启用颜色
		DisableColors:    false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableQuote:     true,
		CallerPrettyfier: callerPretty,
	})

	// 设置日志级别（开发用 Debug，生产可改成 Info/Warn）
	Log.SetLevel(logrus.DebugLevel)
	//log.SetLevel(logrus.InfoLevel)
	//Log.SetFormatter(&logrus.JSONFormatter{})
}

// callerPretty 自定义显示文件:行号
func callerPretty(frame *runtime.Frame) (function string, file string) {
	return "", fmt.Sprintf("%s:%d", path.Base(frame.File), frame.Line)
}
