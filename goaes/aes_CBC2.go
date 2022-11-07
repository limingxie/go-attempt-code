package goaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// CBCEncrypt AES-CBC 加密
// key 必须是 16(AES-128)、24(AES-192) 或 32(AES-256) 字节的 AES 密钥；
// 初始化向量 iv 为随机的 16 位字符串 (必须是16位)，
// 解密需要用到这个相同的 iv，因此将它包含在密文的开头。
func CBCEncrypt(plaintext string, key string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cbc decrypt err:", err)
		}
	}()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	blockSize := len(key)
	padding := blockSize - len(plaintext)%blockSize // 填充字节
	if padding == 0 {
		padding = blockSize
	}

	// 填充 padding 个 byte(padding) 到 plaintext
	plaintext += string(bytes.Repeat([]byte{byte(padding)}, padding))
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil { // 将同时写到 ciphertext 的开头
		return ""
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.StdEncoding.EncodeToString(ciphertext)
}

// CBCDecrypt AES-CBC 解密
func CBCDecrypt(ciphertext string, key string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cbc decrypt err:", err)
		}
	}()

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return ""
	}

	ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	iv := ciphercode[:aes.BlockSize]        // 密文的前 16 个字节为 iv
	ciphercode = ciphercode[aes.BlockSize:] // 正式密文

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)

	plaintext := string(ciphercode) // ↓ 减去 padding
	return plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]
}
func MainGoAesCBC2() {
	key := "123456789abcdefg" // 32位 AES-256

	ciphertext := CBCEncrypt(`{"code":200,"data":{"apts":[]},"message":"","success":true}`, key)

	plaintext := CBCDecrypt(ciphertext, key)

	fmt.Println("ciphertext:", ciphertext)
	fmt.Println("plaintext:", plaintext)

}
