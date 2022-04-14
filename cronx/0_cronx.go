package cronx

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/bcyxy/cronx/common/conf"
	"github.com/bcyxy/cronx/common/gval"
	"github.com/bcyxy/cronx/common/log"
	"github.com/bcyxy/cronx/cronx/controller"
	"github.com/bcyxy/cronx/cronx/worker"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Main() {
	var err error

	// 处理命令行
	handleArgs()

	// 初始化配置
	err = conf.InitConf()
	if err != nil {
		fmt.Printf("Init config failed: %v\n", err)
		return
	}

	// 初始化日志
	l := conf.GetConf("global", "log_level", "INFO")
	logDir := gval.RootDir + "/log"
	os.MkdirAll(logDir, os.ModePerm)
	logPath := logDir + "/cronx.log"
	err = log.Init(l, logPath)
	if err != nil {
		fmt.Printf("Init logger failed: %v\n", err)
		return
	}
	defer log.Defer()

	// 启动controller
	if true {
		err = controller.Start()
		if err != nil {
			log.Error("start_controller_failed. err=%v", err)
			return
		}
		defer worker.Stop()
	}

	// 启动worker
	if true {
		err = worker.Start()
		if err != nil {
			log.Error("start_worker_failed. err=%v", err)
			return
		}
		defer worker.Stop()
	}

	// 等待停止信号
	sCh := make(chan os.Signal, 1)
	signal.Notify(sCh, os.Interrupt, os.Kill)
	<-sCh
}

// 处理命令参数
func handleArgs() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [OPTION]...\n", os.Args[0])
		flag.PrintDefaults()
	}

	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "Only show version.")
	flag.Parse()

	if showVersion {
		fmt.Printf("GitCommitID: %v\n", gval.GitCommitID)
		fmt.Printf("BuildTime:   %v\n", gval.BuildTime)
		os.Exit(0)
	}
	return
}
