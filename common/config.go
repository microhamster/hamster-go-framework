package common

import (
	"context"
	"hamster/core"
	"hamster/core/messager"
	"hamster/log"
)

// 系统配置
type SystemConfig struct {
	Debug        bool                   `mapstructure:"debug"`
	Logout       string                 `mapstructure:"logout"`
	Salt         string                 `mapstructure:"salt"`
	Redis        SystemDataConfig       `mapstructure:"redis"`
	Mysql        SystemDataConfig       `mapstructure:"mysql"`
	Messager     messager.MessageConfig `mapstructure:"messager"`
	ApiConfig    SystemApiConfig        `mapstructure:"api_server"`
	SettleConfig SystemSettleConfig     `mapstructure:"settle_server"`
}

// 系统数据配置
type SystemDataConfig struct {
	Name     string `mapstructure:"name"`
	Endpoint string `mapstructure:"endpoint"`
}

// 系统API配置
type SystemApiConfig struct {
	Pprof                 string `mapstructure:"pprof"`
	Serve                 string `mapstructure:"serve"`
	AllowOrigins          string `mapstructure:"allow_origins"`
	SignaturePrivateKey   string `mapstructure:"signature_private_key"`
	RequestLimitPerMinute int    `mapstructure:"request_limit_per_minute"`
}

// 系统结算配置
type SystemSettleConfig struct {
	Pprof string `mapstructure:"pprof"`
}

// 系统配置
func GetSystemConfig() *SystemConfig {
	config := &SystemConfig{}
	err := core.GetViper().Unmarshal(config)
	if err != nil {
		log.Errorf("failed to get system config: %s", err.Error())
		return nil
	}
	return config
}

// 配置监视
func ConfigMonitor(ctx context.Context, config *SystemConfig) {
	go func() {
		defer core.Recover()
		for {
			select {
			case <-ctx.Done():
				return
			case <-core.ConfigChannel:
				err := core.GetViper().Unmarshal(&config)
				if err != nil {
					log.Errorf("failed to unmarshal system config: %s", err.Error())
				}
			}
		}
	}()
}
