package _chan

import (
	"fmt"
	"testing"
	"time"
)

// struct{}{} 用法
func TestDemo0101(t *testing.T) {
	ch := make(chan struct{})

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("aaa")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("bbb")
	ch <- struct{}{}
}

func TestDemo0102(t *testing.T) {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	// 输出 1，可见上一步的操作不仅仅是个信号，它要占用1个位置
	println(len(ch))
	// 继续添加会报错
	//ch <- struct{}{}
}
