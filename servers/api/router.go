package apiServer

import (
	"hamster/core"
	"hamster/log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/utils"
)

// Web服务
func (s *ApiServer) WebServer() {

	defer core.Recover()

	prefixlog := core.GetFunctionName()

	// 限制上传
	app := fiber.New(fiber.Config{
		BodyLimit: 2 * 1024 * 1024, // 2MB
	})

	// 跨域访问
	app.Use(cors.New(cors.Config{
		AllowOrigins: s.Config.ApiConfig.AllowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// 缓存请求
	app.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Route().Path == "/api/v1/banker/seed"
		},
		Expiration:   1 * time.Second,
		CacheHeader:  "X-Cache",
		CacheControl: false,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
		ExpirationGenerator:  nil,
		StoreResponseHeaders: false,
		Storage:              nil,
		MaxBytes:             0,
		Methods:              []string{fiber.MethodGet, fiber.MethodHead},
	}))

	// 限制频率
	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        s.Config.ApiConfig.RequestLimitPerMinute,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(ResultObject{Code: API_CODE_REQUEST_LIMIT_ERROR, Data: "too many requests"})
		},
	}))

	// 获取系统时间
	app.Get("/api/v1/timestamp", s.GetTimestamp)

	// 启动服务
	log.Infof("%s listen on: %s", prefixlog, s.Config.ApiConfig.Serve)
	err := app.Listen(s.Config.ApiConfig.Serve)
	if err != nil {
		log.Errorf("%s failed to listen on: %s -> %s", prefixlog, s.Config.ApiConfig.Serve, err.Error())
		return
	}

}
