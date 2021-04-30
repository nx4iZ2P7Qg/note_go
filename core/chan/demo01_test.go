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
