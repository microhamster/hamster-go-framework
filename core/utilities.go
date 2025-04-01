package core

import (
	"context"
	"hamster/log"
	"math/big"
	"net/url"
	"runtime"
	"time"

	"github.com/shopspring/decimal"
)

// 线程休眠
func Sleep(ctx context.Context, d time.Duration) {
	timer := time.NewTimer(d)
	select {
	case <-ctx.Done():
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
	}
}

// 获取系统信息
func GetSystemInfo() (memory uint64, goroutine int) {
	var systemMemory runtime.MemStats
	runtime.ReadMemStats(&systemMemory)
	memory = systemMemory.Sys / 1024 / 1024
	goroutine = runtime.NumGoroutine()
	return
}

// 获取函数名
func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	return funcName
}

// 对象转字符串
func ObjectToJson(data interface{}) string {
	text, err := FastJsonMarshalToString(data)
	if err != nil {
		return ""
	}
	return text
}

// 判断是否结束
func ContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// 获取URL参数
func GetUrlPath(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		log.Errorf("failed to get url path: %s", s)
		return nil
	}
	return u
}

// 协程防崩溃
func Recover(args ...any) {
	if r := recover(); r != nil {
		log.Panicf("%s failed to execute goroutine args:%#v r:%v caller_stack:%s", EVENT_APP_PANIC, args, r, GetCallerStackLog())
	}
}

// 代币数量除以精度
func TokenDivDecimal(amount decimal.Decimal, decimals uint8) decimal.Decimal {
	return amount.Div(decimal.NewFromBigInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil), 0))
}

// 代币数量除以精度
func TokenMulDecimal(amount decimal.Decimal, decimals uint8) decimal.Decimal {
	return amount.Mul(decimal.NewFromBigInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil), 0))
}
