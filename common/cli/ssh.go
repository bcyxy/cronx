package cli

import (
	"fmt"

	expect "github.com/google/goexpect"
	//"golang.org/x/crypto/ssh"
)

// SSHCli TELNET客户端
type SSHCli struct {
	opter Operator
	gExp  *expect.GExpect
}

// NewSSHCli 创建telnet客户端
func NewSSHCli() *SSHCli {
	return &SSHCli{
		opter: *NewOperator(),
	}
}

// Login 登入
func (sf *SSHCli) Login(host, user, passwd string) error {
	var err error

	cmd := fmt.Sprintf("ssh %s@%s", user, host)
	sf.gExp, _, err = expect.Spawn(cmd, -1)
	if err != nil {
		return fmt.Errorf("spawn_failed:%v", err)
	}
	return sf.opter.Login(sf.gExp, user, passwd)
}

// Logout 登出
func (sf *SSHCli) Logout() error {
	sf.gExp.Send("exit\n")
	sf.gExp.Send("quie\n")
	return sf.gExp.Close()
}

// ExeCmd 登出
func (sf *SSHCli) ExeCmd(cmd string, turnTime int) (string, error) {
	return sf.opter.Req(sf.gExp, cmd, turnTime)
}
