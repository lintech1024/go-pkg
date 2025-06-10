package main

import (
	"crypto"
	"crypto/ed25519"
	"fmt"
	"log"
)

func main() {
	// 1. 生成密钥对
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatal("生成密钥失败:", err)
	}

	// 2. 签名
	message := []byte("Hello, Ed25519!")
	signature := ed25519.Sign(privKey, message)
	fmt.Printf("签名结果: %x\n", signature)

	// 3. 验证签名
	valid := ed25519.Verify(pubKey, message, signature)
	fmt.Println("签名是否有效:", valid)

	// 4. 带选项的签名验证
	opts := &ed25519.Options{Hash: crypto.SHA512}
	err = ed25519.VerifyWithOptions(pubKey, message, signature, opts)
	if err == nil {
		fmt.Println("带选项验证通过")
	}
}
