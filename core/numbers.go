package core

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/shopspring/decimal"
)

// 数据转字节
func DataToBytes(data interface{}) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, data)
	return bytesBuffer.Bytes()
}

// 整数转浮点
func BigIntToFloat64(amount *big.Int) float64 {
	f, _ := amount.Float64()
	return f
}

// 字符串转浮点
func StringToFloat64(amount string) float64 {
	a, err := decimal.NewFromString(amount)
	if err != nil {
		return 0.0
	}
	result, _ := a.Float64()
	return result
}

// Uint32转小端字节数组
func Uint32ToLittleBytes(number uint32) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, number)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}
