package gval

import (
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	GitCommitID string // 编译的代码版本
	BuildTime   string // 编译时间，格式：2022-02-10 11:10:41
	BuildTs     int64  // 编译时间毫秒时间戳，由BuildTime转换
	StartTs     int64  // 程序启动毫秒时间戳
	RootDir     string // 程序根路径
)

func init() {
	tm, _ := time.ParseInLocation("2006-01-02 15:04:05", BuildTime, time.Local)
	BuildTs = tm.Unix() * 1000

	StartTs = time.Now().UnixNano() / 1e6

	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	RootDir = filepath.Join(filepath.Dir(path), "../")
}
