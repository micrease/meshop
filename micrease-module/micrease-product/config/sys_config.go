package config

import (
	"fmt"
	"github.com/micrease/micrease-core/config"
	"sync"
)

//继承CommonConfig此基础上扩展
type SysConfig struct {
	config.CommonConfig `yaml:",inline"`
}

//自定义配置需要实现Config接口
/*
type CustomConfig struct {
	config.Config
	Name string `yaml:"name "json:"name"`
}
*/

var once sync.Once
var sysConfig *SysConfig

func InitSysConfig() *SysConfig {
	once.Do(func() {
		var err error
		sysConfig, err = config.LoadConfig[SysConfig]()
		fmt.Println(sysConfig)
		if err != nil {
			panic(err)
		}
	})
	return sysConfig
}

func Get() *SysConfig {
	return sysConfig
}
