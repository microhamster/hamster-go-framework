package messager

import (
	"context"
	"hamster/core"
	"hamster/log"
	"time"

	"github.com/go-resty/resty/v2"
)

type TelegramMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type MessageConfig struct {
	Base MessageItemConfig `mapstructure:"base"`
	Warn MessageItemConfig `mapstructure:"warn"`
}

type MessageItemConfig struct {
	Url     string `mapstructure:"url"`
	ChatId  int64  `mapstructure:"chat_id"`
	Keyword string `mapstructure:"keyword"`
}

// 发送消息
func SendMessage(ctx context.Context, chatId int64, url string, keyword string, message string) error {
	content := TelegramMessage{ChatID: chatId, Text: message}

	client := resty.New()
	client.SetHeaders(map[string]string{
		"Content-Type": "application/json",
	})
	// client.SetBaseURL("")
	client.SetRetryCount(1)
	client.SetTimeout(time.Duration(5) * time.Second)
	client.SetRetryWaitTime(time.Duration(0) * time.Second)
	client.AddRetryCondition(func(r *resty.Response, err error) bool {
		if err != nil || r.IsError() {
			log.Errorf("retry to send message: %s -> %s", r.Request.URL, r.String())
			return true
		}
		return false
	})

	body, err := core.FastJsonMarshal(content)
	if err != nil {
		return err
	}

	response, err := client.R().SetBody([]byte(body)).Post(url)
	if err != nil {
		return err
	}

	log.Infof("send telegram message: %s", response.String())

	return nil
}

// 发送通用消息
func SendBaseMessage(ctx context.Context, config *MessageConfig, message string) error {
	return SendMessage(ctx, config.Base.ChatId, config.Base.Url, config.Base.Keyword, message)
}

// 发送异常消息
func SendWarnMessage(ctx context.Context, config *MessageConfig, message string) error {
	return SendMessage(ctx, config.Base.ChatId, config.Warn.Url, config.Warn.Keyword, message)
}
