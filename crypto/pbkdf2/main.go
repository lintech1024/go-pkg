package main

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"fmt"
)

func main() {
	password := "my-secure-password"
	salt := []byte("random-salt") // 实际应使用 crypto/rand 生成随机盐
	keyLength := 32               // AES-256 需要 32 字节
	iter := 4096                  // 推荐至少 10,000 次迭代

	key, err := pbkdf2.Key(sha256.New, password, salt, iter, keyLength)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Derived Key: %x\n", key)
}
