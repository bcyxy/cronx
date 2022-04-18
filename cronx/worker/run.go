package worker

import (
	"time"

	"github.com/bcyxy/cronx/common/log"
	"github.com/bcyxy/cronx/cronx/worker/jobconf"
	"github.com/bcyxy/cronx/cronx/worker/jobmgr"
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

		// 与外部同步配置
		isChg, err := jobconf.LoadConf()
		if err != nil {
			log.Warn("worker:load_job_conf_failed:%v", err)
			continue
		}
		if !isChg {
			continue
		}
		log.Info("worker:job_conf_changed")

		// 生成作业表
		var jobs interface{}
		jobs, err = jobconf.GenJobs()
		if err != nil {
			log.Warn("worker:xxx_failed:%v", err)
			continue
		}

		// 更新任务
		err = jobmgr.UpdateAll(jobs)
		if err != nil {
			log.Warn("worker:xxx_failed:%v", err)
			continue
		}
	}
}
