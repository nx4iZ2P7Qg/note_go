package main

import (
	"fmt"
	"strings"
	"time"
)

// IPAddr ip数组
type IPAddr [4]byte

// 实现 String() string
// 删除本方法后再运行，会以 [4]byte 的格式输出
func (ip *IPAddr) String() string {
	var s []string
	for _, b := range ip {
		// 返回数值形式的串
		s = append(s, fmt.Sprintf("%d", b))
	}
	return strings.Join(s, ".")
}

// 类型与其方法，实现 String() string 接口
func demo0101() {
	host := &IPAddr{127, 0, 0, 1}
	fmt.Print(host)
}

// MyError 自定义错误
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// 类型与其方法，实现 Error() string 接口
func demo0102() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
