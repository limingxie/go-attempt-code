package gomethod

import "fmt"

type User_r struct {
	id   int
	name string
}

func (self *User_r) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

func ReceiverTest() {
	u := User_r{1, "Tom"}
	u.Test()

	mValue := u.Test
	mValue() // 隐式传递 receiver

	mExpression := (*User_r).Test
	mExpression(&u) // 显式传递 receiver
	/*
	   0xc42000a060, &{1 Tom}
	   0xc42000a060, &{1 Tom}
	   0xc42000a060, &{1 Tom}
	*/
}

func (self User_r) Test1() {
	fmt.Println(self)
}

func ReceiverTest1() {
	u := User_r{1, "Tom"}
	mValue := u.Test1 // 立即复制 receiver，因为不是指针类型，不受后续修改影响。

	u.id, u.name = 2, "Jack"
	u.Test1() //{2 Jack}

	mValue() //{1 Tom}
}

func (self *User_r) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", self, self)
}

func (self User_r) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &self, self)
}

func ReceiverTest2() {
	u := User_r{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u) //User: 0xc00000c048, {1 Tom}

	mv := User_r.TestValue
	mv(u) //TestValue: 0xc00000c078, {1 Tom}

	mp := (*User_r).TestPointer
	mp(&u) //TestPointer: 0xc00000c048, &{1 Tom}

	mp2 := (*User_r).TestValue // *User 方法集包含 TestValue。签名变为 func TestValue(self *User)。实际依然是 receiver value copy。
	mp2(&u)                    //TestValue: 0xc00000c0c0, {1 Tom}
}

func MainReceiverTest() {
	ReceiverTest2()
}
