package cli

import "context"

/* 说明
- CLI: command line interface，命令行交互界面。主要用于人与机器交互，而非给程序调用；
- 特点
	- 全双工：在协议层，客户端在发送消息的同时，服务的也可以同时发消息；
	- 边界模糊：交互类似两个人聊天，没有确定的规则，只能通过对方语义判断对方是否结束发言；
	- 性能低：因为没有确定的结束标识，所以收到结束字符后必须稍等一下；
- 实现：程序调研CLI功能需要模拟人的操作，以telnet为例：
	[root@yq01-sys-netadmin02.yq01 ~]# telnet 192.168.172.88
	Escape character is '^]'.

	login: zhangsan             # 看到login提示，输入用户名
	Password:                   # 看到Password提示，输入密码
	Last login: Fri Nov 26 08:51:41 from 192.168.172.6

	[standalone: master] >      # 判断登录成功，看到>，发一个\n测试一下。
	[standalone: master] > show version
	Product name:      MLNX-OS
	。。。
	[standalone: master] >      # 判断命令回显是否结束，一般看到字符>会等一秒，以确定没有后续数据。
	[standalone: master] > quit # 有始有终
	Connection closed by foreign host.
	[root@yq01-sys-netadmin02.yq01 ~]# */

// Cli 客户端
type Cli interface {
	// Login 登录
	Login(ctx context.Context, host, user, passwd string) error

	// Logout 登出
	Logout() error

	// ExeCmd 执行命令。turnTime: 自动翻页次数，小于0时一直翻页
	ExeCmd(ctx context.Context, cmd string, turnTime int) (stdout, stderr string, err error)
}
