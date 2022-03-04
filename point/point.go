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

func MainPoint() {

	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	// s1 = append(s1, 1)
	// fmt.Printf("%p\n", &s1)

	s2 := append(s1, 2)
	fmt.Printf("%p\n", s2)

	// s1 := make([]int, 0, 5)
	// fmt.Printf("s1: %p, %v\n", &s1, s1)

	// // s1 = append(s1, 1)
	// // fmt.Printf("%p\n", &s1)

	// s2 := append(s1, 2)
	// fmt.Printf("s2: %p, %v\n", s2, s2)

	// s1 = append(s1, 1)
	// fmt.Printf("s1: %p, %v\n", &s1, s1)
	// fmt.Printf("s2: %p, %v\n", s2, s2)

	// s2[0] = 10

	// fmt.Println(s1, s2)

	// s1 = append(s1, 3)
	// fmt.Printf("%p, %p\n", &s1, &s2)

	// fmt.Println(s1, s2)

	// fmt.Printf("%p, %p\n", &s1[0], &s2[0])

	// s3 := &s1
	// fmt.Printf("%p\n", &s3)

}
