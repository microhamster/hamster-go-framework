package apiServer

import (
	"context"
	"fmt"

	"hamster/common"
	"hamster/core"
	"hamster/log"
	"time"

	"github.com/redis/go-redis/v9"
)

type ApiServer struct {
	Ctx       context.Context      // 上下文
	CtxCancel context.CancelFunc   // 全局控制
	Runtime   int64                // 运行时间
	Redis     *redis.Client        // 缓存连接
	Config    *common.SystemConfig // 系统配置
}

// 获取服务
func GetApiServer() *ApiServer {
	ctx, cancel := context.WithCancel(context.Background())
	data := &ApiServer{
		Ctx:       ctx,
		CtxCancel: cancel,
		Runtime:   time.Now().Unix(),
		Config:    common.GetSystemConfig(),
	}
	return data
}

// 创建服务
func NewApiServer() core.Runable {
	return GetApiServer()
}

// 初始化参数
func (s *ApiServer) Init() bool {

	// 初始化配置
	common.ConfigMonitor(s.Ctx, s.Config)

	// 初始化监控
	if len(s.Config.ApiConfig.Pprof) > 0 && s.Config.Debug {
		go core.PprofMonitor(s.Config.ApiConfig.Pprof)
	}

	// 初始化日志
	log.Init(s.Config.Debug, s.Config.Logout)

	// 初始化redis
	s.Redis = core.OpenRedis(s.Ctx, s.Config.Redis.Endpoint)
	if s.Redis == nil {
		return false
	}

	return true
}

// 启动服务
func (s *ApiServer) Start() {

	if !s.Init() {
		return
	}

	// 启动Web服务
	go s.WebServer()

	log.Info(fmt.Sprintf("%s %s", core.GetFunctionName(), core.EVENT_APP_START))

}

// 停止服务
func (s *ApiServer) Stop() {

	s.CtxCancel()

	log.Info(fmt.Sprintf("%s %s", core.GetFunctionName(), core.EVENT_APP_STOP))
}

// 退出服务
func (s *ApiServer) Done() <-chan struct{} {
	return s.Ctx.Done()
}
