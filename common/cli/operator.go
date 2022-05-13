package cli

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	expect "github.com/google/goexpect"
)

//TODO 可能有认证时间较长的情况

var regxAny = regexp.MustCompile(`.`)

// Operator 操作员
type Operator struct {
	RegUserSign   string // 正则:输入用户名提示符
	RegPasswdSign string // 正则:输入密码提示符
	RegEscapeSign string // 正则:回显结束提示符
	RegMoreSign   string // 正则:分段显示提示符

	regxEsc  *regexp.Regexp // 回显结束提示符
	regxMore *regexp.Regexp
}

// NewOperator 新建CLI客户端
func NewOperator() *Operator {
	return &Operator{
		RegUserSign:   `([Uu]sername|[Ll]ogin): ?$`,
		RegPasswdSign: `[Pp]assword: ?$`,
		RegEscapeSign: "[>#\\$][\x00-\x09\x7F\x0B-\x1F ]?$",
		//RegMoreSign:   `(--.?[Mm]ore.?-+|lines \d+-\d+).{0,20}$`,
		RegMoreSign: "\x1B\\[\\d+mlines \\d+-\\d+.{0,15}\x1B\\[\\d+m\x1B\\[K$", // Mella
	}
}

// Login 登录操作
func (sf *Operator) Login(gExp *expect.GExpect, user, passwd string) error {
	// 编译匹配正则
	regxUser, err := regexp.Compile(sf.RegUserSign)
	if err != nil {
		return fmt.Errorf("compile_login_reg_failed:%v", err)
	}
	regxPasswd, err := regexp.Compile(sf.RegPasswdSign)
	if err != nil {
		return fmt.Errorf("compile_password_reg_failed:%v", err)
	}
	sf.regxEsc, err = regexp.Compile(sf.RegEscapeSign)
	if err != nil {
		return fmt.Errorf("compile_password_reg_failed:%v", err)
	}
	sf.regxMore, err = regexp.Compile(sf.RegMoreSign)
	if err != nil {
		return fmt.Errorf("compile_more_reg_failed:%v", err)
	}

	// 交互
	var buf string
	for {
		s, _, err := gExp.Expect(regxAny, 10*time.Second)
		if err != nil {
			return err
		}
		buf += s

		if sf.regxEsc.MatchString(buf) {
			s, _, err = gExp.Expect(regxAny, 500*time.Millisecond)
			if err != nil && strings.Contains(err.Error(), "timer expired") {
				return nil
			}
			buf += s
		} else if regxUser.MatchString(buf) {
			time.Sleep(100 * time.Millisecond)
			gExp.Send(user + "\n") //TODO 是否所有系统都是\n
			buf = ""
		} else if regxPasswd.MatchString(buf) {
			time.Sleep(100 * time.Millisecond)
			gExp.Send(passwd + "\n")
			buf = ""
		}
	}
}

// Req 请求
func (sf *Operator) Req(gExp *expect.GExpect, cmd string, turnTime int) (string, error) {
	time.Sleep(100 * time.Millisecond)
	err := gExp.Send(cmd + "\n")
	if err != nil {
		return "", fmt.Errorf("send_cmd_failed:%v", err)
	}

	var buf string
	turnCnt := 0
	for {
		s, _, err := gExp.Expect(regxAny, 10*time.Second)
		if err != nil {
			return "", err
		}
		buf += s
		fmt.Printf("####%s\n", s)
		fmt.Printf("====%v\n", []byte(s))

		if sf.regxEsc.MatchString(buf) {
			/*s, _, err = gExp.Expect(regxAny, 500*time.Millisecond)
			if err != nil && strings.Contains(err.Error(), "timer expired") {
				return buf, nil
			}
			buf += s*/
			return buf, nil
		} else if sf.regxMore.MatchString(buf) {
			if turnTime < 0 || turnCnt < turnTime {
				fmt.Println("#send space.")
				time.Sleep(100 * time.Millisecond)
				buf = sf.regxMore.ReplaceAllString(buf, "")
				gExp.Send(" ")
				turnCnt++
			} else {
				fmt.Println("#send q.")
				time.Sleep(100 * time.Millisecond)
				buf = sf.regxMore.ReplaceAllString(buf, "")
				gExp.Send("q")
			}

		}
	}
}
