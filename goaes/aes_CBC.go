package goaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

/*
	AES-CBC 加密
	text String 待加密的明文
	key String 秘钥
	iv String iv值
	return Base64, err
*/
func EncryptAESCBC(text, key, iv string) (result string, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16、24、32")
	}

	if len(key) != len(iv) {
		return "", errors.New("IV length must equal block size")
	}

	encrypted, err := CBCEncrypter([]byte(text), []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

/*
	AES-CBC 解密
	encrypter base64 待解密的密文
	key string 秘钥
	iv string iv值
	return String, err
*/
func DecryptAESCBC(encrypter, key, iv string) (result string, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("请确保的key的长度是 16、24、32 中一个。")
	}

	if len(key) != len(iv) {
		return "", errors.New("IV length must equal block size")
	}

	encryptByte, err := base64.StdEncoding.DecodeString(encrypter)
	if err != nil {
		return "", err
	}

	decrypted, err := CBCDecrypter(encryptByte, []byte(key), []byte(iv))
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

/*
	CBC 加密
	text 待加密的明文
	key 秘钥
*/
func CBCEncrypter(text []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 填充
	paddText := PKCS7Padding(text, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, iv)

	// 加密
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	// 返回密文
	return result, nil
}

/*
	CBC 解密
	encrypter 待解密的密文
	key 秘钥
*/
func CBCDecrypter(encrypter []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(encrypter))
	blockMode.CryptBlocks(result, encrypter)
	// 去除填充
	result = UnPKCS7Padding(result)
	return result, err
}

func MainGoAesCBC() {
	key := "123456789abcdefg"
	iv := "0123456789abcdef"
	text := "aes_CBC加密测试"

	// 加密 []byte
	encrypted, _ := CBCEncrypter([]byte(text), []byte(key), []byte(iv))
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))

	// 解密 []byte
	decrypted, _ := CBCDecrypter([]byte(encrypted), []byte(key), []byte(iv))
	fmt.Println(string(decrypted))

	// 加密 string
	encryptedStr, _ := EncryptAESCBC(text, key, iv)
	fmt.Println(encryptedStr)
	// 解密 string
	decryptStr, _ := DecryptAESCBC(encryptedStr, key, iv)
	fmt.Println(decryptStr)
}
