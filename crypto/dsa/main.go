package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

func main() {
	// 1. 生成 DSA 参数
	params := new(dsa.Parameters)
	err := dsa.GenerateParameters(params, rand.Reader, dsa.L1024N160)
	if err != nil {
		panic(err)
	}

	// 2. 生成密钥对
	priv := new(dsa.PrivateKey)
	priv.PublicKey.Parameters = *params
	err = dsa.GenerateKey(priv, rand.Reader)
	if err != nil {
		panic(err)
	}

	// 3. 签名
	hash := []byte("Hello, DSA!")
	r, s, err := dsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		panic(err)
	}
	fmt.Printf("签名结果: r=%x, s=%x\n", r.Bytes(), s.Bytes())

	// 4. 验证签名
	valid := dsa.Verify(&priv.PublicKey, hash, r, s)
	fmt.Println("签名是否有效:", valid)
}
