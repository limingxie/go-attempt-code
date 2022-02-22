package linkedlist

import (
	"fmt"
	"strings"
)

type HashMap struct {
	BucketMap []HashNode
	Size      int
}

type HashNode struct {
	Key   string
	Value int
	Next  *HashNode
}

func GetHashCodeIndex(key string) int {
	return len(key) % 10
}

func (h *HashMap) Get(key string) int {
	bucketIndex := GetHashCodeIndex(key)
	hashNode := &h.BucketMap[bucketIndex]
	if len(hashNode.Key) == 0 && hashNode.Value == 0 && hashNode.Next == nil {
		return -1
	} else {
		for strings.Compare(hashNode.Key, key) != 0 {
			hashNode = hashNode.Next
		}
		if hashNode == nil {
			return -1
		} else {
			return hashNode.Value
		}
	}
}

func (h *HashMap) Add(key string, value int) {
	bucketIndex := GetHashCodeIndex(key)
	hashNode := &h.BucketMap[bucketIndex]

	if len(hashNode.Key) == 0 && hashNode.Value == 0 && hashNode.Next == nil {
		h.BucketMap[bucketIndex] = HashNode{Key: key, Value: value}
		h.Size++
	} else {
		for hashNode.Next != nil {
			hashNode = hashNode.Next
		}
		hashNode.Next = &HashNode{Key: key, Value: value}
		h.Size++
	}
}

func (h *HashMap) Remove(key string) {
	bucketIndex := GetHashCodeIndex(key)
	hashNode := &h.BucketMap[bucketIndex]

	if len(hashNode.Key) == 0 && hashNode.Value == 0 && hashNode.Next == nil {
		return
	} else {
		parentNode := hashNode
		for strings.Compare(hashNode.Key, key) != 0 {
			parentNode = hashNode
			hashNode = hashNode.Next
		}
		if hashNode == nil {
			return
		} else if parentNode == hashNode {
			h.BucketMap[bucketIndex] = *hashNode.Next
			h.Size--
		} else {
			parentNode.Next = hashNode.Next
			h.Size--
		}
	}
}

func (h *HashMap) IsEmpty() bool {
	return h.Size == 0
}

func (h *HashNode) Print() {
	if h == nil || h.Value == 0 {
		return
	} else {
		fmt.Print(h.Value, " ")
		h.Next.Print()
	}
}

func MainHashTable() {
	bucketMap := make([]HashNode, 10)
	hashMap := HashMap{BucketMap: bucketMap}
	hashMap.Add("a", 1)
	hashMap.Add("ab", 2)
	hashMap.Add("abc", 3)
	hashMap.Add("abcd", 4)
	hashMap.Add("abcde", 5)
	hashMap.Add("1234567890ab", 12)
	hashMap.Add("1234567890abc", 13)
	hashMap.Add("12345678901234567890abc", 113)
	hashMap.Add("1234567890xabcde", 15)

	fmt.Println("-------------Size--------------")
	fmt.Println("Size:", hashMap.Size)

	fmt.Println("-------------Get--------------")
	fmt.Println(hashMap.Get("1234567890abc"))

	fmt.Println("-------------Add--------------")
	fmt.Print("Bucket[3]: ")
	hashMap.BucketMap[3].Print()
	fmt.Println("")

	hashMap.Add("123456789012345678901234567890abc", 1113)
	fmt.Print("Bucket[3]: ")
	hashMap.BucketMap[3].Print()
	fmt.Println("")
	fmt.Println("Size:", hashMap.Size)

	fmt.Println("-------------Remove--------------")
	hashMap.Remove("abc")
	fmt.Print("Bucket[3]: ")
	hashMap.BucketMap[3].Print()
	fmt.Println("")
	fmt.Println("Size:", hashMap.Size)
}
