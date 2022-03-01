package slice

import (
	"unsafe"
)

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

//如果切片的容量小于 1024 个元素，于是扩容的时候就翻倍增加容量。上面那个例子也验证了这一情况，总容量从原来的4个翻倍到现在的8个。
//一旦元素个数超过 1024 个元素，那么增长因子就变成 1.25 ，即每次增加原来容量的四分之一。
