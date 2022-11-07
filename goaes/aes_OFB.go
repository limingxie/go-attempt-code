package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

// 加密 分别返回 hex格式和base64 结果
func AesEncryptByOFB(data, key string) (string, string) {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// 转成byte
	dataByte := []byte(data)
	keyByte := []byte(key)
	// 创建block
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		panic(fmt.Sprintf("NewCipher error:%s", err))
	}
	blockSize := block.BlockSize()
	// 创建偏移量iv,取秘钥前16个字符
	iv := []byte(key[:blockSize])
	//fmt.Printf("iv = %s \n",iv)
	// 补码
	padding := PKCS7Padding(dataByte, blockSize)
	// 加密
	stream := cipher.NewOFB(block, iv)
	// 定义保存结果变量
	out := make([]byte, len(padding))
	stream.XORKeyStream(out, padding)
	// 处理加密结果
	hexRes := fmt.Sprintf("%x", out)
	base64Res := base64.StdEncoding.EncodeToString(out)
	return hexRes, base64Res
}

// 解密
func AesDecryptByOFB(dataBase64, key string) string {
	// 判断key长度
	keyLenMap := map[int]struct{}{16: {}, 24: {}, 32: {}}
	if _, ok := keyLenMap[len(key)]; !ok {
		panic("key长度必须是 16、24、32 其中一个")
	}
	// dataBase64转成[]byte
	decodeStringByte, err := base64.StdEncoding.DecodeString(dataBase64)
	if err != nil {
		panic(fmt.Sprintf("base64 DecodeString error: %s", err))
	}
	// 创建block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(fmt.Sprintf("NewCipher error: %s", err))
	}
	blockSize := block.BlockSize()
	// 创建偏移量iv,取秘钥前16个字符
	iv := []byte(key[:blockSize])
	// 创建Stream
	stream := cipher.NewOFB(block, iv)
	// 声明变量
	out := make([]byte, len(decodeStringByte))
	// 解密
	stream.XORKeyStream(out, decodeStringByte)
	// 解密加密结果并返回
	return string(PKCS7UNPadding(out))
}

func MainGoAesOFB() {
	key := "123456789abcdefg"
	data := "aes_CBC加密测试"

	hex, base64 := AesEncryptByOFB(data, key)
	fmt.Printf("加密key: %v \n", key)
	fmt.Printf("加密数据: %v \n", data)

	fmt.Printf("加密结果(hex): %v \n", hex)
	fmt.Printf("加密结果(base64): %v \n", base64)

	res := AesDecryptByOFB(base64, key)
	fmt.Printf("解密结果: %v \n", res)
}
