package gomethod

import "fmt"

/***********************************/

//结构体
type User struct {
	Name  string
	Email string
}

//方法
func (u *User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}
func GomenthodTest() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	/*
		golang : golang@golang.com
		go : go@go.com
	*/
}

/***********************************/

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
	fmt.Printf("Pointer: %p\n", self)
}

func gomethodTest1() {
	d := Data{}
	p := &d
	fmt.Printf("Data: %p\n", p)

	d.ValueTest()   // ValueTest(d)
	d.PointerTest() // PointerTest(&d)

	p.ValueTest()   // ValueTest(*p)
	p.PointerTest() // PointerTest(p)

	/*
			Data: 0xc42007c008
		    Value: 0xc42007c018
		    Pointer: 0xc42007c008
		    Value: 0xc42007c020
		    Pointer: 0xc42007c008
	*/
}

/***********************************/
//1.普通函数
//接收值类型参数的函数
func valueIntTest(a int) int {
	return a + 10
}

//接收指针类型参数的函数
func pointerIntTest(a *int) int {
	return *a + 10
}

func structTestValue() {
	a := 2
	fmt.Println("valueIntTest:", valueIntTest(a)) //valueIntTest: 12
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println("valueIntTest:", valueIntTest(&a))
	//compile error: cannot use &a (type *int) as type int in function argument

	b := 5
	fmt.Println("pointerIntTest:", pointerIntTest(&b)) //pointerIntTest: 15
	//同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	//fmt.Println("pointerIntTest:", pointerIntTest(b))
	//compile error:cannot use b (type int) as type *int in function argument
}

//2.方法
type PersonD struct {
	id   int
	name string
}

//接收者为值类型
func (p PersonD) valueShowName() {
	fmt.Println(p.name)
}

//接收者为指针类型
func (p *PersonD) pointShowName() {
	fmt.Println(p.name)
}

func structTestFunc() {
	//值类型调用方法
	personValue := PersonD{101, "hello world"}
	personValue.valueShowName() //hello world
	personValue.pointShowName() //hello world

	//指针类型调用方法
	personPointer := &PersonD{102, "hello golang"}
	personPointer.valueShowName() //hello golang
	personPointer.pointShowName() //hello golang

	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func GomethodTest2() {
	structTestValue()
	structTestFunc()
}

/***********************************/

func Maingomethod() {
	Gotest1()
}

/***********************************/

type User_a struct {
	id   int
	name string
}

type Manager struct {
	User_a
	title string
	name  string
}

func (self *User_a) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func Gotest1() {
	m := Manager{User_a{1, "Tom"}, "Administrator", "Administrator"}

	fmt.Println(m.ToString())        //Manager: 0xc420074180, &{{1 Tom} Administrator}
	fmt.Println(m.User_a.ToString()) //User: 0xc420074180, &{1 Tom}

	fmt.Printf("%p, %v\n", &m, m)               //0xc000106040, {{1 Tom} Administrator Administrator}
	fmt.Printf("%p, %v\n", &m.User_a, m.User_a) //0xc000106040, {1 Tom}
}
