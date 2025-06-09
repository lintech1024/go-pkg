package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// PKCS7 填充函数
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7 去填充函数
func pkcs7Unpadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func main() {
	key := []byte("example key 1234") // 16 字节密钥
	plaintext := []byte("AES加密测试数据")

	// 创建 AES 分组加密器
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// CBC 模式需要 16 字节的 IV
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	// 填充明文以满足块大小要求
	paddedText := pkcs7Padding(plaintext, block.BlockSize())

	// 加密
	ciphertext := make([]byte, len(paddedText))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)
	fmt.Printf("加密结果: %x\n", ciphertext)

	// 解密
	decrypted := make([]byte, len(ciphertext))
	decryptMode := cipher.NewCBCDecrypter(block, iv)
	decryptMode.CryptBlocks(decrypted, ciphertext)

	// 去除填充
	decrypted = pkcs7Unpadding(decrypted)
	fmt.Printf("解密结果: %s\n", decrypted)
}
