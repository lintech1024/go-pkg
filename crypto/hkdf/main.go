package main

import (
	"crypto/hkdf"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	// 主密钥（如 TLS 预主密钥）
	masterSecret := []byte("master-secret")

	// 盐值（推荐随机生成）
	salt := []byte("random-salt")

	// 提取伪随机密钥
	prk, err := hkdf.Extract(sha256.New, masterSecret, salt)
	if err != nil {
		log.Fatal("Extract 失败:", err)
	}
	fmt.Printf("伪随机密钥: %x\n", prk)

	// 扩展生成多个子密钥
	aesKey, _ := hkdf.Expand(sha256.New, prk, "aes-key", 32)
	hmacKey, _ := hkdf.Expand(sha256.New, prk, "hmac-key", 32)
	fmt.Printf("AES 密钥: %x\n", aesKey)
	fmt.Printf("HMAC 密钥: %x\n", hmacKey)

	// 使用 Key 函数一步生成密钥
	finalKey, _ := hkdf.Key(sha256.New, masterSecret, salt, "final-key", 16)
	fmt.Printf("最终密钥: %x\n", finalKey)
}
