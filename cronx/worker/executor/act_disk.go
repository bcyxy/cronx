package executor

import (
	"context"
	"encoding/json"
	"fmt"
)

func init() {
	RegisterAct(&actDisk{})
}

type actDisk struct{}

type sshObj struct {
	Host   string `json:"host"`
	User   string `json:"user"`
	Passwd string `json:"password"`
}

type sshParam struct {
	Cmd string `json:"cmd"`
}

func (sf *actDisk) GetName() string {
	return "DISK"
}

// Do 执行
// return ["/,/dev/vda1,82437508,14609796", "/run,tmpfs,2022508,824"]
func (sf *actDisk) Do(ctx context.Context, objStr, actStr string) (ret []string, err error) {
	// 解析参数
	var obj sshObj
	err = json.Unmarshal([]byte(objStr), &obj)
	if err != nil {
		err = fmt.Errorf("unmarshal_obj_failed:%v", err)
		return
	}
	var act sshParam
	err = json.Unmarshal([]byte(actStr), &act)
	if err != nil {
		err = fmt.Errorf("unmarshal_act_failed:%v", err)
		return
	}

	// sshExec(obj, "df")
	fmt.Println(obj, act)

	return
}

func sshExec(host, cmd string) {}
