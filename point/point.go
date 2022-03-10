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
	*c = *c * 3

	fmt.Println(&a)
	fmt.Println(c) //都是同一个值地址

	fmt.Println(a)  //int和float 乘法是一样的计算方式所以不出问题
	fmt.Println(*c) //int的值用float的方式读取，所以出问题。

	*c = *c - 2

	fmt.Println(a)  //int和float 乘法是一样的计算方式所以不出问题
	fmt.Println(*c) //int的值用float的方式读取，所以出问题。
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

func uintptrPointerTest() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + 9*unsafe.Sizeof(a[0]))

	// b是 unsafe.Pointer 所以可转任意指针，转成(*int)指针后在取值
	fmt.Printf("b: %v, unsafe.Sizeof(a[0]): %d\n", *(*int)(b), unsafe.Sizeof(a[0])) //b: 9, unsafe.Sizeof(a[0]): 8

	c := unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + uintptr(16)) //int是8位长度 所以16 等于 16/8 挪动了2位，所以下面结果是2
	fmt.Printf("c: %v\n", *(*int)(c))                              //c: 2

	user := user{id: 1, age: 10, name: "user1"}
	namePointer := unsafe.Pointer(uintptr(unsafe.Pointer(&user)) + unsafe.Offsetof(user.name))

	//这也一样 name是 unsafe.Pointer 所以可转任意指针，转成(*string)指针后在取值
	fmt.Printf("name: %v\n", *(*string)(namePointer)) //name: user1
}

func slicePointTest() {
	//因slice的结构是 => |ptr|len|cap
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8))) //挪一个位置是Len
	fmt.Println(Len, len(s))                                                    // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16))) //挪二个位置是CAP
	fmt.Println(Cap, cap(s))                                                     // 20 20

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	//因map结构中第一个是元素个数，所以可以转成len
	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

func MainPoint() {
	slicePointTest()
}
