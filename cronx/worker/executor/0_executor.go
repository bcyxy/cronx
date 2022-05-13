package executor

import (
	"context"
	"fmt"
	"time"

	"github.com/bcyxy/cronx/common/gstruct"
)

// ActIf 动作接口
type ActIf interface {
	Do(ctx context.Context, obj, act string) ([]string, error)
	GetName() string
}

var actMap = make(map[string]ActIf)

// Do 执行任务
func Do(ctx context.Context, task *gstruct.Task) {
	task.StartTs = time.Now().UnixMilli()
	defer func() {
		task.EndTs = time.Now().UnixMilli()
	}()

	// 选act
	act, ok := actMap[task.Type]
	if !ok {
		task.Err = fmt.Errorf("unknow_type")
		return
	}

	// 执行
	func() {
		defer func() {
			if err := recover(); err != nil {
				task.Err = fmt.Errorf("act_panic:%v", err)
			}
		}()
		task.Ret, task.Err = act.Do(ctx, task.Obj, task.Act)
	}()
	return
}

// RegisterAct 注册动作
func RegisterAct(act ActIf) {
	actMap[act.GetName()] = act
}
