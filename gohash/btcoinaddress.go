package gohash

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func HASH256(text string, isHex bool) (string, error) {
	sha256Instance := sha256.New()
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return "", err
		}
		sha256Instance.Write(arr)
	} else {
		sha256Instance.Write([]byte(text))
	}
	ciphertext := sha256Instance.Sum(nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func HASH256HexString(text string) (string, error) {
	arr, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}
	sha256Instance := sha256.New()
	sha256Instance.Write(arr)
	ciphertext := sha256Instance.Sum(nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func RipMD160(text string, isHex bool) (string, error) {
	ripemd160Instance := ripemd160.New()
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return "", err
		}
		ripemd160Instance.Write(arr)
	} else {
		ripemd160Instance.Write([]byte(text))
	}
	hashBytes := ripemd160Instance.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	return hashString, nil
}

//获取校验码
func GetCheckCode(hexString string) string {
	sha256DoubleString, _ := Hash256DoubleString(hexString, true)

	rs := []rune(sha256DoubleString)
	checkCode := string(rs[:8])
	return checkCode
}

func GenerateAddress(pubKey string, nettype int) (string, error) {
	//1.判断公钥的有效性
	if len(pubKey) != 130 && len(pubKey) != 66 {
		return "", errors.New("公钥输入的长度不合法！")
	}
	//2.计算公钥sha256
	hashString, err := HASH256(pubKey, true)
	if err != nil {
		return "", err
	}

	//3.ripemd160
	ripemd160String, err := RipMD160(hashString, true)
	if err != nil {
		return "", err
	}

	//4.添加网络id号，比特币主网0x00 testnet 0x6f
	prefix := ""
	switch nettype {
	case 0:
		prefix = "00"
	case 1:
		prefix = "6f"
	case 2:
		prefix = "34"
	}

	versionString := prefix + ripemd160String

	//5.计算hash 2次
	sha256DoubleString, err := Hash256DoubleString(versionString, true)
	if err != nil {
		return "", err
	}

	//6.获取校验码
	rs := []rune(sha256DoubleString)
	checknum := string(rs[:8])

	//7.形成16进制比特币地址
	addrHex := versionString + checknum

	//8.对16进制地址进行base58编码
	arr, err := hex.DecodeString(addrHex)
	if err != nil {
		return "", err
	}

	addrBase58 := Base58Encode(arr)
	result := fmt.Sprintf("%s \n", addrBase58)

	return result, nil
}

func MainBitCoin() {
	pubKey := "0417a4a4037defe1c30da0c3ed0ca096ae99e846a10c594b0224cb4c54cc0d087028e27b218c422d095a11c8b9a34e2bfdb2f112311adea31245e190e3775cc5b7"
	address, _ := GenerateAddress(pubKey, 0)
	fmt.Println(address) //1Lr12b7kWNXbRiwPEGq65Y19NuKYXyPJiK
}
