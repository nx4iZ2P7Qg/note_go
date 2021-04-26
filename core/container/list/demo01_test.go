package list

import (
	"container/list"
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	// 创建双向链表
	r := list.New()
	r.PushBack("b1")
	r.PushFront("f1")
	r.PushBack("b2")
	r.PushFront("f2")

	for e := r.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
