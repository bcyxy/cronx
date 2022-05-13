package gstruct

import (
	"time"

	"github.com/bcyxy/cronx/common/gfunc"
)

// Task 任务：对一个目标做一套动作
type Task struct {
	ID       uint64   // 任务ID
	Type     string   // 任务类型
	Obj      string   // 操作对象
	Act      string   // 动作
	Ret      []string // 任务输出
	Err      error    // 执行报错
	CreateTs int64    // 任务创建时间
	StartTs  int64    // 任务开始执行时间
	EndTs    int64    // 任务执行结束时间
}

func NewTask() *Task {
	return &Task{
		ID:       uint64(gfunc.GetUniqID()),
		Ret:      make([]string, 0),
		CreateTs: time.Now().UnixMilli(),
	}
}
