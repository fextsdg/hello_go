package main

//接口练习

import (
	"fmt"
)

// 新建接口
type A interface {
	sayHello()
	sayHello1()
}

//结构体

type B struct {
	name string
	age  int
}

func (b B) sayHello() { //接口实现
	fmt.Printf("%s,%d岁，向你打招呼！\n", b.name, b.age)
}

type C struct {
	B     //继承
	class string
}

func (c C) sayHello() {
	fmt.Printf("类型%s:%s,%d岁，向你打招呼！\n", c.class, c.name, c.age)
}
func (c C) sayHello1() {
	fmt.Printf("类型%s:%s,%d岁，向你打招呼1！\n", c.class, c.name, c.age)
}

// 只有完全实现了接口内的函数的结构体才可以作为下面函数的参数传入
func greet(a A) {
	a.sayHello1()
	a.sayHello()
}
func main() {
	var b B = B{name: "小王", age: 18}
	b.sayHello()
	c := C{
		B: B{
			name: "小李",
			age:  18,
		},
		class: "C",
	}
	c.sayHello()
	c.B.sayHello()
	c.sayHello1()
	greet(c)
}
