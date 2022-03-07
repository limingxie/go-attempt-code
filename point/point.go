package point

import (
	"fmt"
	"reflect"
	"unsafe"
)

//int 和 float64 都指向同一个地址
func PointerTypeTest() {
	var a int = 1
	// var b *int = &a

	var c *float64 = (*float64)(unsafe.Pointer(&a))
	*c = *c * 100

	fmt.Println(&a)
	fmt.Println(c)

	fmt.Println(a)
	fmt.Println(*c)
}

func PointerTypeTest1() {
	var a float64 = 1

	var b *float64 = &a

	var c *int = (*int)(unsafe.Pointer(b))
	fmt.Println(a)
	*c = (*c) * 2

	fmt.Println(a)
	fmt.Println(*c)
}

func PointerTypeTest2() {
	var a int = 1
	b := (*int)(unsafe.Pointer(&a))

	fmt.Println(&a)
	fmt.Println(b)

}

type user struct {
	id   int
	age  int
	name string
}

func PointerStructTest() {
	u := new(user)
	fmt.Printf("%+v\n", *u)

	pAge := (*int)(unsafe.Pointer(u)) // 第一个不用计算指针偏移量
	*pAge = 20

	pName := (*string)(unsafe.Pointer(u))
	*pName = "张二"

	fmt.Printf("%+v\n", *u)

	pName1 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.name)))
	*pName1 = "张三"

	pAge1 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge1 = 30

	fmt.Printf("%+v\n", *u)
}

func PointTest() {
	v1 := uint(12)
	v2 := int(13)

	fmt.Println(reflect.TypeOf(v1)) //uint
	fmt.Println(reflect.TypeOf(v2)) //int

	fmt.Println(reflect.TypeOf(&v1)) //*uint
	fmt.Println(reflect.TypeOf(&v2)) //*int

	p := &v1
	p = (*uint)(unsafe.Pointer(&v2)) //使用unsafe.Pointer进行类型的转换

	fmt.Println(reflect.TypeOf(p)) // *unit
	fmt.Println(*p)                //13
}

func uintptrTest() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		a[i] = i
	}
	fmt.Println(a)
	// [0 1 2 3 4 5 6 7 8 9]

	// 取slice的最后的一个元素
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + 9*unsafe.Sizeof(a[0]))
	// 等价于unsafe.Pointer(&b[9])
	fmt.Println(*(*int)(end))
	// 9

}

type s1 struct {
	id   int
	name string
}

type s2 struct {
	field1 *[5]byte
	filed2 int
}

func pointTest1() {
	b := s1{name: "123"}
	var j s2
	j = *(*s2)(unsafe.Pointer(&b))
	fmt.Println(j)
}

func MainPoint() {
	pointTest1()

}
