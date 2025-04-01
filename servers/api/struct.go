package apiServer

import "github.com/gofiber/fiber/v2"

// API返回对象
type ResultObject struct {
	Code int         `json:"code"` // 成功为0，失败为自定义Code
	Data interface{} `json:"data"` // 任意数据对象
}

// API分页返回对象
type ResultPaging struct {
	Code      int         `json:"code"`       // 成功为0，失败为自定义Code
	Data      interface{} `json:"data"`       // 任意数据对象
	PageIndex int         `json:"page_index"` // 当前页码
	PageSize  int         `json:"page_size"`  // 每页数量
	PageCount int         `json:"page_count"` // 总页数量
}

// 返回成功消息
func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(ResultObject{
		Code: API_CODE_SUCCESS,
		Data: data,
	})
}

// 返回分页消息
func Paging(c *fiber.Ctx, data ResultPaging) error {
	return c.JSON(data)
}

// 返回错误消息
func Error(c *fiber.Ctx, code int, data string) error {
	return c.JSON(ResultObject{
		Code: code,
		Data: data,
	})
}
