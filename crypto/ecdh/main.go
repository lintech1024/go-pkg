package main

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func main() {
	// 选择曲线（推荐 X25519）
	curve := ecdh.X25519()

	// 生成私钥和公钥
	privA, pubA := generateKeyPair(curve)
	privB, pubB := generateKeyPair(curve)

	// 计算共享密钥
	sharedA, err := privA.ECDH(pubB)
	if err != nil {
		panic(err)
	}
	sharedB, err := privB.ECDH(pubA)
	if err != nil {
		panic(err)
	}

	// 验证共享密钥一致性
	fmt.Printf("共享密钥 A: %x\n", sharedA)
	fmt.Printf("共享密钥 B: %x\n", sharedB)
	fmt.Println("密钥是否一致:", string(sharedA) == string(sharedB))
}

// 生成密钥对
func generateKeyPair(curve ecdh.Curve) (*ecdh.PrivateKey, *ecdh.PublicKey) {
	priv, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	return priv, priv.PublicKey()
}
