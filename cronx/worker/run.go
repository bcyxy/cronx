package worker

import (
	"fmt"
	"time"
)

func run() {
	defer stopSyn.Done()

	tkr := time.NewTicker(time.Second)
	for {
		select {
		case <-stopCh:
			return
		case <-tkr.C:
		}

		fmt.Println("worker running")

		// 与外部同步配置
		// 生成作业表
	}
}
