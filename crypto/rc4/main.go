package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	key := []byte("1234567890") // 密钥长度需在 1~256 字节之间
	plaintext := []byte("Hello, RC4!")

	// 创建 RC4 加密器
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// 加密
	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("Encrypted: %x\n", ciphertext)

	// 解密（重新初始化 Cipher）
	cipher, _ = rc4.NewCipher(key)
	decrypted := make([]byte, len(ciphertext))
	cipher.XORKeyStream(decrypted, ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
