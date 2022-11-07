package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//GCM 全称为Galois/Counter Mode，可以看出 G 是指 GMAC，C 是指 CTR。它在 CTR 加密的基础上增加 GMAC 的特性，解决了 CTR 不能对加密消息进行完整性校验的问题。
// 加密(GCM 不需要补码)
func AesEncryptByGCM(data, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	// 生成随机因子(这里固定取密钥指定位数)
	//nonce := make([]byte, gcm.NonceSize())
	//if _,err := io.ReadFull(rand.Reader,nonce); err != nil {
	//	panic(fmt.Sprintf("make rand nonce error:%s", err))
	//}
	nonceStr := key[:gcm.NonceSize()]
	nonce := []byte(nonceStr)
	fmt.Printf("nonceStr = %v \n", nonceStr)
	seal := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.StdEncoding.EncodeToString(seal)
}

// 解密(GCM 不需要解码)
func AesDecryptByGCM(data, key string) string {
	// 反解base64
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(fmt.Sprintf("base64 DecodeString error:%s", err))
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(fmt.Sprintf("NewGCM error:%s", err))
	}
	nonceSize := gcm.NonceSize()
	if len(dataByte) < nonceSize {
		panic("dataByte to short")
	}
	nonce, ciphertext := dataByte[:nonceSize], dataByte[nonceSize:]
	open, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(fmt.Sprintf("gcm Open error:%s", err))
	}
	return string(open)
}

func MainGoAesGCM() {
	key := "123456789abcdefg"
	data := "aes_CBC加密测试"
	// 加密
	gcm := AesEncryptByGCM(data, key)
	fmt.Printf("密钥key: %s \n", key)
	fmt.Printf("加密数据: %s \n", data)

	fmt.Printf("加密结果: %s \n", gcm)
	// 解密
	byGCM := AesDecryptByGCM(gcm, key)
	fmt.Printf("解密结果: %s \n", byGCM)
}
