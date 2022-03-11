package slice

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

//如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
//一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。

func MainSliceCode() {
	var a []int = nil
	var b *slice = nil

	fmt.Printf("%p, %p, %v\n", a, &a, &a)
	fmt.Printf("%p, %p, %v\n", b, &b, &b)
}

func SliceTest20() {
	arrayA := [2]int{100, 200}
	arrayB := arrayA[:]
	fmt.Printf("arrayA : %p, %p, %p, %v\n", &arrayA, &arrayA[0], &arrayA[1], arrayA) //arrayA : 0xc0000b0010, 0xc0000b0010, 0xc0000b0018, [100 200]
	fmt.Printf("arrayB : %p, %p, %p, %v\n", arrayB, &arrayB[0], &arrayB[1], arrayB)  //arrayB : 0xc0000b0010, 0xc0000b0010, 0xc0000b0018, [100 200]
}

func SliceTest21() {
	s := make([]int, 5)
	ptr := unsafe.Pointer(&s[0])
	fmt.Printf(" %p, %p, %v\n", s, &s[0], s)    //0xc0000a8030, 0xc0000a8030, [0 0 0 0 0]
	fmt.Printf(" %p, %p, %v\n", ptr, &ptr, ptr) //0xc0000a8030, 0xc0000aa020, 0xc0000a8030
}
