package core

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hamster/log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forgoer/openssl"
	"golang.org/x/crypto/sha3"
)

// Sha256加密
func Sha256ToBytes(data string) []byte {
	b32 := sha256.Sum256([]byte(data))
	return b32[:]
}

// 生成Sha256字节数组
func Keccak256ToBytes(paramsters ...[]byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	for _, input := range paramsters {
		hasher.Write(input)
	}
	return hasher.Sum(nil)
}

// 密钥加密
func Encrypt(salt string, data string) string {
	privateKey, err := openssl.AesECBEncrypt([]byte(data), Sha256ToBytes(fmt.Sprintf("%s.%s", HAMSTER_ENCRYPT_SALT, salt)), openssl.PKCS7_PADDING)
	if err != nil {
		log.Errorf("failed to encrypt data: %s - %s", data, err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(privateKey)
}

// 密钥解密
func Decrypt(salt string, data string) string {
	base64Key, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Errorf("failed to decrypt data: %s - %s", data, err.Error())
		return ""
	}
	secretKey, err := openssl.AesECBDecrypt(base64Key, Sha256ToBytes(fmt.Sprintf("%s.%s", HAMSTER_ENCRYPT_SALT, salt)), openssl.PKCS7_PADDING)
	if err != nil {
		log.Errorf("failed to decrypt data: %s - %s", data, err.Error())
		return ""
	}
	return string(secretKey)
}

// 获取私钥对象
func GetAesPrivateKey(salt string, data string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(Decrypt(salt, data))
	if err != nil {
		log.Errorf("failed to get ecdsa private key: %s", err.Error())
		return nil
	}
	return privateKey
}

// 获取私钥对象
func GetEcdsaPrivateKey(secret string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(secret)
	if err != nil {
		log.Errorf("failed to get ecdsa private key: %s", err.Error())
		return nil
	}
	return privateKey
}
