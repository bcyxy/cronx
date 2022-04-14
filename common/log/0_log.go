package log

import (
	"os"
	"path"
	"strings"
	"time"

	"github.com/hyahm/golog"
)

var (
	fileLog *golog.Log
	toFile  bool
)

// Init 初始化日志模块
func Init(dir, fName string) error {
	// 日志级别
	l := "INFO"
	switch strings.ToUpper(l) {
	case "DEBUG":
		golog.Level = golog.DEBUG
	case "INFO":
		golog.Level = golog.INFO
	case "WARN":
		golog.Level = golog.WARN
	case "ERROR":
		golog.Level = golog.ERROR
	default:
		golog.Level = golog.All
	}

	// 控制台打印
	golog.Format = `{{ .Ctime }}|{{ .Level }}|{{.Msg}}`

	// 输出到文件
	os.MkdirAll(dir, os.ModePerm)
	fileLog = golog.NewLog(path.Join(dir, fName), 0, true, int(time.Hour*24*10))
	fileLog.Format = golog.Format
	toFile = true

	return nil
}

// InitTmp 临时初始化，用于测试等
func InitTmp() {
	golog.Level = golog.All
	golog.Format = `{{ .Ctime }}|{{ .Level }}|{{.Msg}}`
	toFile = false
}

// Defer 模块结束时做的操作
func Defer() {
	golog.Sync()
}

// Debug debug级日志
func Debug(format string, args ...interface{}) {
	golog.Debugf(format, args...)
	if toFile {
		fileLog.Debugf(format, args...)
	}
}

// Info info级日志
func Info(format string, args ...interface{}) {
	golog.Infof(format, args...)
	if toFile {
		fileLog.Infof(format, args...)
	}
}

// Warn warn级日志
func Warn(format string, args ...interface{}) {
	golog.Warnf(format, args...)
	if toFile {
		fileLog.Warnf(format, args...)
	}
}

// Error error级日志
func Error(format string, args ...interface{}) {
	golog.Errorf(format, args...)
	if toFile {
		fileLog.Errorf(format, args...)
	}
}
