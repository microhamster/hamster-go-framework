package core

import (
	"math/big"
	"unsafe"

	"github.com/shopspring/decimal"
)

// 字符串转字节
func StringTobytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// 字节转字符串
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 字符串转大整数
func StringToBigInt(amount string) *big.Int {
	d, err := decimal.NewFromString(amount)
	if err != nil {
		return big.NewInt(0)
	} else {
		return d.BigInt()
	}
}

// 字符串转大浮点
func StringToBigFloat(amount string) *big.Float {
	number := new(big.Float)
	result, ok := number.SetString(amount)
	if !ok {
		return big.NewFloat(0)
	}
	return result
}

// hex串转大整数
func HexToBigNumber(amount string) *big.Int {
	number := new(big.Int)
	result, ok := number.SetString(amount, 16)
	if !ok {
		return big.NewInt(0)
	}
	return result
}
