package gohash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func getMerkleRootHash() string {
	//先声明4个交易单的hash
	coinBase := "1cd8352f53e0d8e8914e7c4e2509c5648288390fc74c207c1b5482919a783def"
	tx1 := "4006d7cad120309f6bab70e45319ee8b519eb63c4e91bb7972ce5a052471b848"
	tx2 := "4e8782354afc43152b48177af49eb971acd1e209384bf816c9380dc47e60fff2"
	tx3 := "3f4bb8454d44b4bdd48cad4a17c32757f8b432469d815a35a4ae4a789cdcf2bb"

	//大小端转换处理
	r_coinbase := ReverseStr2ByteStr(coinBase)
	r_tx1 := ReverseStr2ByteStr(tx1)
	r_tx2 := ReverseStr2ByteStr(tx2)
	r_tx3 := ReverseStr2ByteStr(tx3)

	//两两计算hash

	//计算hash先转arr
	r_coinbase_tx1_arr, err := hex.DecodeString(r_coinbase + r_tx1)
	if err != nil {
		fmt.Println(err)
	}

	//每次hash计算都是2次
	res_c1 := sha256.Sum256(r_coinbase_tx1_arr)
	res_c1 = sha256.Sum256(res_c1[:])

	r_tx2_tx3_arr, err := hex.DecodeString(r_tx2 + r_tx3)
	if err != nil {
		fmt.Println(err)
	}

	//每次hash计算都是2次
	res_34 := sha256.Sum256(r_tx2_tx3_arr)
	res_34 = sha256.Sum256(res_34[:])

	//字节数组转换成字符相加
	r_res_c123_str := hex.EncodeToString(res_c1[:]) + hex.EncodeToString(res_34[:])

	//两两再一次计算hash
	r_res_c123_arr, err := hex.DecodeString(r_res_c123_str)
	if err != nil {
		fmt.Println(err)
	}

	//2次hash计算
	result := sha256.Sum256(r_res_c123_arr)
	result = sha256.Sum256(result[:])

	//最后一个hash再做一次大小端转换。
	merkleRootHash := ReverseStr2ByteStr(hex.EncodeToString(result[:]))

	return merkleRootHash
}

func ReverseStr2ByteStr(str string) string {
	result := ""
	for i := len(str) / 2; i > 0; i-- {
		result += str[i*2-2 : i*2]
	}
	return result
}

func MainGoMerkleRootHash() {
	fmt.Println(getMerkleRootHash())
}
