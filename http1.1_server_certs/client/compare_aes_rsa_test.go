package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"io"
	"testing"
)


// func main(){}

// TLSでは一度だけ公開鍵で厳重に通信し、その後は共通鍵で高速に暗号化通信をしている
// 公開鍵は厳重な分計算量が大きい
// go test -bench . でベンチマークを取ってみる
// 公開鍵の用意
func prepareRSA() (sourceData, label []byte, privateKey *rsa.PrivateKey) {
	sourceData = make([]byte, 128)
	label = []byte("")
	io.ReadFull(rand.Reader, sourceData)
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	return
}

// 公開鍵: 暗号化
func BenchmarkRSAEnctyption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA()
	publicKey := &privateKey.PublicKey
	md5hash := md5.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)
	}
}

// 公開鍵: 復号化
func BenchmarkRSADecryption(b *testing.B) {
	sourceData, label, privateKey := prepareRSA()
	publicKey := &privateKey.PublicKey
	md5hash := md5.New()
	encrypted, _ := rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.DecryptOAEP(md5hash, rand.Reader, privateKey, encrypted, label)
	}
}

// 共通鍵を用意
func PrepareAES() (sourceData, nonce []byte, gcm cipher.AEAD) {
	sourceData = make([]byte, 128)
	io.ReadFull(rand.Reader, sourceData)
	key := make([]byte, 32)
	io.ReadFull(rand.Reader, sourceData)
	nonce = make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)
	block, _ := aes.NewCipher(key)
	gcm, _ = cipher.NewGCM(block)
	return
}

// 共通鍵: 暗号化
func BenchmarkAESEncryption(b *testing.B) {
	sourceData, nonce, gcm := PrepareAES()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Open(nil, nonce, sourceData, nil)
	}
}

// 共通鍵: 復号化
func BenchmarkAESDncryption(b *testing.B) {
	sourceData, nonce, gcm := PrepareAES()
	encrypted := gcm.Seal(nil, nonce, sourceData, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gcm.Open(nil, nonce, encrypted, nil)
	}
}
