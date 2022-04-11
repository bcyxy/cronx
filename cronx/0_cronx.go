package cronx

import (
	"flag"
	"fmt"
	"os"

	"github.com/bcyxy/cronx/common/gval"
)

func Main() {}

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
