package executor_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/bcyxy/cronx/common/gstruct"
	"github.com/bcyxy/cronx/cronx/worker/executor"
)

func TestExector(t *testing.T) {
	task := gstruct.NewTask()
	task.Type = "DISK"
	task.Obj = `{"host":"127.0.0.1", "user":"yxy", "password":"666666"}`
	task.Act = `{"cmd": "df"}`
	executor.Do(context.TODO(), task)
	fmt.Printf("%+v\n", *task)
}
