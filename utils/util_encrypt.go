package utils

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	mRand "math/rand"
	"strings"
	"time"
)

func RsaEncrypt(publicKey, origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(privateKey, cipherText []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	var priKey interface{}
	var err error
	if block.Type == "RSA PRIVATE KEY" {
		//解析PKCS1格式的私钥
		priKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else if block.Type == "PRIVATE KEY" {
		//解析PKCS8格式的私钥
		priKey, err = x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), cipherText)
}

func AesEncryptECB(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS7Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted, nil
}

func AesDecryptECB(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return PKCS7UnPadding(decrypted), nil
}

func GenRsaKey(bits int) (err error, privateKey string, publicKey []byte) {
	// 生成私钥文件
	originKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}

	derStream := x509.MarshalPKCS1PrivateKey(originKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	stringPrivateKey := string(pem.EncodeToMemory(priBlock))
	slicePrivateKey := strings.Split(stringPrivateKey, "\n")
	for i, v := range slicePrivateKey {
		if i == 0 || i > len(slicePrivateKey)-3 {
			continue
		}

		privateKey = privateKey + v
	}

	// 生成公钥文件
	originPublicKey := &originKey.PublicKey
	derPkIx, err := x509.MarshalPKIXPublicKey(originPublicKey)
	if err != nil {
		return
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkIx,
	}
	publicKey = pem.EncodeToMemory(publicBlock)

	return err, privateKey, publicKey
}

func GetRandomString(num int) []byte {
	letterByte := []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789`)

	b := make([]byte, num)
	r := mRand.New(mRand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letterByte[r.Intn(len(letterByte))]
	}

	return b
}
