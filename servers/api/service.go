package apiServer

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// 获取系统时间 /api/v1/timestamp
func (s *ApiServer) GetTimestamp(c *fiber.Ctx) error {
	return Success(c, time.Now().Unix())
}
