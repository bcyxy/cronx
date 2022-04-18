package executor

import (
	"context"
	"fmt"
)

type ActIf interface {
	Do(ctx context.Context, obj, param string) (map[string]string, error)
	GetName() string
}

var actMap = make(map[string]ActIf)

func Do(ctx context.Context, actType, objKey, actParam string) (ret map[string]string, err error) {
	act, ok := actMap[actType]
	if !ok {
		return nil, fmt.Errorf("unknow_act_type")
	}

	func() {
		defer func() {
			if err := recover(); err != nil {
				err = fmt.Errorf("act_panic:%v", err)
			}
		}()
		ret, err = act.Do(ctx, objKey, actParam)
	}()
	return
}
