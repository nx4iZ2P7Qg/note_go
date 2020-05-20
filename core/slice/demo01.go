package main

import (
	"bytes"
	"fmt"
	"log"
)

func demo0101() {
	s1 := []byte{1, 2, 3}
	s2 := []byte{1, 2, 3}

	// slice 没有 == 操作，除了 == nil
	fmt.Printf("bytes.Equal() = %v", bytes.Equal(s1, s2))
	// 除 []byte 其他的 Equal 需要自己实现
}

func demo0102() {
	var s []string
	log.Printf("s == nil, %v", s == nil)
	s = append(s, "a")
	log.Printf("nil slice append")
}
