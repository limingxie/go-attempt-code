package gohash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"

	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
)

func HASH(text string, hashType string, isHex bool) (string, error) {
	var hashInstance hash.Hash
	switch hashType {
	case "md4":
		hashInstance = md4.New()
	case "md5":
		hashInstance = md5.New()
	case "sha1":
		hashInstance = sha1.New()
	case "ripemd160":
		hashInstance = ripemd160.New()
	case "sha256":
		hashInstance = sha256.New()
	case "sha512":
		hashInstance = sha512.New()
	}
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return "", err
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}

	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", bytes), nil
}

func HashMain() {
	hash256Str, err := HASH("123456", "sha256", false)
	if err == nil {
		fmt.Println(hash256Str)
	} else {
		fmt.Println(err)
	}
}

func Hash256Double(text string, isHex bool) ([]byte, error) {
	hashInstance := sha256.New()
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return nil, err
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}

	bytes := hashInstance.Sum(nil)
	hashInstance.Reset()
	hashInstance.Write(bytes)
	bytes = hashInstance.Sum(nil)
	return bytes, nil
}

func Hash256DoubleString(text string, isHex bool) (string, error) {
	bytes, err := Hash256Double(text, isHex)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

func MD4(text string, isHex bool) (string, error) {
	var hashInstance hash.Hash
	hashInstance = md4.New()
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return "", err
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}

	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", bytes), nil
}

func MD5(text string, isHex bool) (string, error) {
	var hashInstance hash.Hash
	hashInstance = md5.New()
	if isHex {
		arr, err := hex.DecodeString(text)
		if err != nil {
			return "", err
		}
		hashInstance.Write(arr)
	} else {
		hashInstance.Write([]byte(text))
	}

	bytes := hashInstance.Sum(nil)
	return fmt.Sprintf("%x", bytes), nil
}
