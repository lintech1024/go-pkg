package main

import (
	"crypto/des"
	"fmt"
	"log"
)

func main() {
	// DES 加密示例
	key := []byte("12345678") // 8 字节
	block, err := des.NewCipher(key)
	if err != nil {
		log.Fatal("DES 密钥错误:", err)
	}

	plaintext := []byte("HelloDES!")
	ciphertext := make([]byte, len(plaintext))

	// 加密
	block.Encrypt(ciphertext, plaintext)
	fmt.Printf("DES Encrypted: %x\n", ciphertext)

	// 解密
	decrypted := make([]byte, len(ciphertext))
	block.Decrypt(decrypted, ciphertext)
	fmt.Printf("DES Decrypted: %s\n", decrypted)

	// 3DES 加密示例
	tripleKey := []byte("1234567890123456") // 16 字节
	tripleBlock, err := des.NewTripleDESCipher(tripleKey)
	if err != nil {
		log.Fatal("3DES 密钥错误:", err)
	}

	tripleCiphertext := make([]byte, len(plaintext))
	tripleBlock.Encrypt(tripleCiphertext, plaintext)
	fmt.Printf("3DES Encrypted: %x\n", tripleCiphertext)
}
