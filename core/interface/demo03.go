package main

import (
	"fmt"
)

type interface1 interface {
	getName() string
}
type interface2 interface {
	printName()
}

type Person struct {
	name string
}

// 注意 Person 同时实现了 interface1 interface2
func (p Person) getName() string {
	return p.name
}

func (p Person) printName() {
	fmt.Println(p.name)
}

// 断言
func demo0301() {
	var t interface1
	t = Person{"xiaohua"}
	check(t)
}

func check(t interface1) {
	// T 非接口
	if f, ok1 := t.(Person); ok1 {
		fmt.Printf("%T %s\n", f, f.getName())
	}
	// T 接口，检查t是否实现 interface2，如果是，t 的类型，值不变，声明类型变更为 interface2
	if t, ok2 := t.(interface2); ok2 {
		fmt.Printf("%T\n", t)
		check2(t)
	}
}

func check2(t interface2) {
	t.printName()
}
