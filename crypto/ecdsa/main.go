package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	// 1. 生成 ECDSA 密钥对
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// 2. 签名
	message := []byte("Hello, ECDSA!")
	hash := sha256.Sum256(message)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("签名结果: r=%x, s=%x\n", r.Bytes(), s.Bytes())

	// 3. 验证签名
	valid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Println("签名是否有效:", valid)

	// 4. ASN.1 编码签名
	asn1Sig, _ := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	validASN1 := ecdsa.VerifyASN1(publicKey, hash[:], asn1Sig)
	fmt.Println("ASN1签名是否有效:", validASN1)
}
