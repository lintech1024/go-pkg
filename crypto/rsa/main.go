package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func main() {
	// 1. 生成 RSA 密钥对
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("生成密钥失败:", err)
	}
	pubKey := &privKey.PublicKey

	// 2. OAEP 加密
	msg := []byte("RSA-OAEP 加密数据")
	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, msg, nil)
	if err != nil {
		log.Fatal("加密失败:", err)
	}
	fmt.Printf("密文: %x\n", ciphertext)

	// 3. OAEP 解密
	plaintext, err := privKey.Decrypt(nil, ciphertext, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatal("解密失败:", err)
	}
	fmt.Printf("明文: %s\n", plaintext)

	// 4. PSS 签名
	hash := sha256.Sum256([]byte("data to sign"))
	signature, err := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, hash[:], nil)
	if err != nil {
		log.Fatal("签名失败:", err)
	}

	// 5. PSS 签名验证
	err = rsa.VerifyPSS(pubKey, crypto.SHA256, hash[:], signature, nil)
	if err != nil {
		log.Fatal("签名验证失败:", err)
	}
	fmt.Println("签名验证通过")
}
