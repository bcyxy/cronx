/* controller
# 描述
- controller负责将任务对象分配给worker；
- 系统中只有一个controller，由多个候选者竞选产生；

# 功能
- 发下任务对象变化，分配；
- 发现worker实例变化，调整任务执行者；*/

package controller

import (
	"fmt"
	"sync"
	"time"
)

var (
	stopCh  chan struct{}
	stopSyn sync.WaitGroup
)

func Start() error {
	stopCh = make(chan struct{})
	stopSyn.Add(1)

	go func() {
		defer stopSyn.Done()

		tkr := time.NewTicker(time.Second)
		for {
			select {
			case <-stopCh:
				return
			case <-tkr.C:
			}
			fmt.Println("controller running")
			// 竞选
		}
	}()

	return nil
}

func Stop() error {
	close(stopCh)
	stopSyn.Wait()
	return nil
}
