package core

import (
	"bytes"
	"fmt"
	"hamster/log"
	"os"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var ConfigChannel chan string

// 配置文件
func GetViper() *viper.Viper {
	return viper.GetViper()
}

// 初始化配置变更信号量
func InitConfigChan() {
	ConfigChannel = make(chan string, 1)
}

// 发送main.yaml配置变更信号量
func SendConfigChange(msg string) {
	if ConfigChannel != nil {
		ConfigChannel <- msg
	}
}

var v *viper.Viper

// ReplaceViper 配置热更，替换为新配置对象
func ReplaceViper(newConf *viper.Viper) {
	v = newConf
}

// ReloadConfig 重新加载配置
func ReloadConfig(cfgType string) error {
	// 回调时，日志已初始化过
	confJSON, err := os.ReadFile(viper.ConfigFileUsed())
	if err != nil {
		log.Warnf("event:%s msg:read config file err %s", "event_conf_err", err.Error())
		return err
	}
	v := viper.New()
	v.SetConfigType(cfgType)
	err = v.ReadConfig(bytes.NewBuffer(confJSON))
	if err != nil {
		// EventConfErr
		log.Errorf("event:%s msg:viper parse config err %s", "event_conf_err", err.Error())
		return err
	}
	ReplaceViper(v) // 新配置替换
	// UpConf
	log.Infof("event:%s msg:update config success", "up_conf")
	return nil
}

var CfgFile string

// initConfig reads in config file and ENV variables if set.
func InitConfig() {

	command := ""
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	InitConfigChan()

	viper.SetConfigType("yaml")
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("main.yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

		v := viper.New()
		data, err := os.ReadFile(viper.ConfigFileUsed())
		if err != nil {
			fmt.Println("read config file err:", err)
			os.Exit(0)
		}
		v.SetConfigType("yaml")
		v.ReadConfig(bytes.NewBuffer(data))
		ReplaceViper(v) // 全新只读实例

		// 监控配置更新
		// 配置文件保存操作会触发回调
		// 注意：ubuntu下软链接的配置文件无法自动触发热更新
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Infof("event:conf_reload")
			// OnConfigChange可能有个bug, 修改配置文件后, 会触发两次。
			// 2次调用对本程序运行无影响，不特殊处理
			ReloadConfig("yaml")
			// 发送配置更新通知
			go SendConfigChange(command)

		})
		viper.WatchConfig()

	} else {
		var skip bool
		if len(os.Args) >= 2 {
			if os.Args[1] == "version" || os.Args[1] == "help" {
				skip = true
			}
		}
		if !skip {
			fmt.Println("config file err:", err)
			os.Exit(1)
		}
	}
}

// 调用栈日志
func GetCallerStackLog() (stacktrace string) {
	for i := 1; ; i++ {
		_, f, l, got := runtime.Caller(i)
		if !got {
			break
		}
		stacktrace += fmt.Sprintf("%s:%d\n", f, l)
	}
	return
}
