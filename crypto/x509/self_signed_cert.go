package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func main() {
	// 生成 2048 位 RSA 密钥
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 证书模板
	template := x509.Certificate{
		SerialNumber: big.NewInt(1), // 序列号
		Subject: pkix.Name{
			Organization: []string{"Example Org"},
			CommonName:   "example.com",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1, 0, 0), // 有效期 1 年

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, // 服务器认证用途
	}
	// 使用模板和私钥生成自签名证书
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &privKey.PublicKey, privKey)
	// 保存证书到 PEM 文件
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 保存私钥到 PEM 文件
	keyOut, _ := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privKey)})
	keyOut.Close()
}
