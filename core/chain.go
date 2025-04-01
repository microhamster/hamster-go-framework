package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
)

// 转换WEI至GWEI
func WeiToGWEI(value *big.Int) float64 {
	amount, _ := decimal.NewFromBigInt(value, 0).Div(decimal.NewFromBigInt(big.NewInt(1), 9)).Float64()
	return amount
}

// 转换WEI至ETH
func WeiToEth(value *big.Int) float64 {
	amount, _ := decimal.NewFromBigInt(value, 0).Div(decimal.NewFromBigInt(big.NewInt(1), 18)).Float64()
	return amount
}

// 转换ETH至WEI
func EthToWei(value float64) *big.Int {
	amount := decimal.NewFromFloat(value).Mul(decimal.NewFromBigInt(big.NewInt(1), 18)).BigInt()
	return amount
}

// 转换ETH至GWEI
func EthToGWei(value float64) *big.Int {
	amount := decimal.NewFromFloat(value).Mul(decimal.NewFromBigInt(big.NewInt(1), 9)).BigInt()
	return amount
}

// 转换GWEI至ETH
func GWeiToEth(value *big.Int) float64 {
	amount, _ := decimal.NewFromBigInt(value, 0).Div(decimal.NewFromBigInt(big.NewInt(1), 9)).Float64()
	return amount
}

// 验证签名消息
func VerifySignature(address string, message string, signature string) bool {
	if address == signature {
		return true
	}
	signer := hexutil.MustDecode(signature)
	if signer[crypto.RecoveryIDOffset] == 27 || signer[crypto.RecoveryIDOffset] == 28 {
		signer[crypto.RecoveryIDOffset] -= 27
	}
	recovered, err := crypto.SigToPub(accounts.TextHash([]byte(message)), signer)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return address == recoveredAddr.Hex()
}
