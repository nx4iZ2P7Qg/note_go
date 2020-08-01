package main

import (
	"errors"
	"fmt"
)

func demo0101() {
	// 2个方法都生成 error
	// 本 wrapper 方法更常用，因为带格式输出
	err := fmt.Errorf("abc")
	fmt.Printf("%T\n", err)
	// 一般不用
	err = errors.New("def")
	fmt.Printf("%T\n", err)
}
