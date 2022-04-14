package conf

import (
	"strings"

	"github.com/Unknwon/goconfig"
	"github.com/bcyxy/cronx/common/gval"
)

var cfg *goconfig.ConfigFile

// InitConf 初始化模块
// TODO 目录参数
func InitConf() (err error) {
	cfg, err = goconfig.LoadConfigFile(gval.RootDir + "/conf/cronx.ini")
	return
}

// GetConf 获取配置
func GetConf(section, key, defVal string) (val string) {
	var err error
	val, err = cfg.GetValue(section, key)
	if err != nil {
		val = defVal
	}
	return
}

// GetConfSections 获取Section列表
func GetConfSections(prefix string) (sections []string) {
	sections = []string{}
	allSections := cfg.GetSectionList()
	for _, section := range allSections {
		if !strings.HasPrefix(section, prefix) {
			continue
		}
		sections = append(sections, section)
	}
	return
}
