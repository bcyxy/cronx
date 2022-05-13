package gfunc

import (
	"sync/atomic"
	"time"
)

var cnter int64

// 获取唯一id
func GetUniqID() int64 {
	return time.Now().UnixMilli()*1000000 + atomic.AddInt64(&cnter, 1)%1000000
}
