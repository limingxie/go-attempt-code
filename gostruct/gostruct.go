package gostruct

import "fmt"

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func StructColumnPoint() {
	n := test{
		1, 2, 3, 4,
	}
	fmt.Printf("n.a %p\n", &n.a) //n.a 0xc000012088
	fmt.Printf("n.b %p\n", &n.b) //n.b 0xc000012089
	fmt.Printf("n.c %p\n", &n.c) //n.c 0xc00001208a
	fmt.Printf("n.d %p\n", &n.d) //n.d 0xc00001208b
}

type Student struct {
	Id   int    `json:"id"` //结构第字段标签(Tag)  通过指定tag实现json序列化该字段时的key
	Name string //json序列化是默认使用字段名作为key
	age  int    //私有不能被json包访问
}

func TestMapStruct() {
	m := make(map[string]*Student)
	stus := []Student{
		{Name: "pprof.cn", age: 18},
		{Name: "测试", age: 23},
		{Name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}

	/*
		pprof.cn => 博客
		测试 => 博客
		博客 => 博客
	*/
}

//NewStudent 构造函数 							=> 这时函数
func NewStudent(name string, age int) *Student {
	return &Student{
		Name: name,
		age:  age,
	}
}

func (s Student) SetAge1(newAge int) {
	s.age = newAge // 起不到作用
}

//接收者(Receiver) 是指针的时候才能起到作用		=> 这时接受者的方法
func (s *Student) SetAge2(newAge int) {
	s.age = newAge
}

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
//注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

/*------------------继承测试-------------------*/
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

func (a *Animal) run() {
	fmt.Printf("%s会跑！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func (a *Dog) run() {
	fmt.Printf("%s会跑而且很快！\n", a.name)
}

func TestInheritance() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang()       //乐乐会汪汪汪~
	d1.move()       //乐乐会动！ yin Dog 没有 move() 方法所以 d1.Animal.move() == d1.move()
	d1.run()        //乐乐会跑而且很快！
	d1.Animal.run() //乐乐会跑！
}

/*------------------继承测试-------------------*/

func MainGoStruct() {
	TestInheritance()
}
