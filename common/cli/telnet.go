package cli

import (
	"fmt"
	"time"

	expect "github.com/google/goexpect"
)

// TelnetCli TELNET客户端
type TelnetCli struct {
	opter Operator
	gExp  *expect.GExpect
}

// NewTelnetCli 创建telnet客户端
func NewTelnetCli() *TelnetCli {
	return &TelnetCli{
		opter: *NewOperator(),
	}
}

// Login 登入
func (sf *TelnetCli) Login(host, user, passwd string) error {
	var err error

	sf.gExp, _, err = expect.Spawn("telnet "+host, -1)
	if err != nil {
		return fmt.Errorf("spawn_failed:%v", err)
	}
	return sf.opter.Login(sf.gExp, user, passwd)
}

// Logout 登出
func (sf *TelnetCli) Logout() error {
	time.Sleep(100 * time.Millisecond)
	sf.gExp.Send("exit\n")
	sf.gExp.Send("quie\n")
	return sf.gExp.Close()
}

// ExeCmd 登出
func (sf *TelnetCli) ExeCmd(cmd string, turnTime int) (string, error) {
	return sf.opter.Req(sf.gExp, cmd, turnTime)
}
