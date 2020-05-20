package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

// Builder
func demo0101() {
	var b strings.Builder
	for i := 3; i >= 1; i-- {
		// 往io.Writer中写数据
		fmt.Fprintf(&b, "%d...", i)
	}
	b.WriteString("ignition")
	fmt.Println(b.String())
}

func demo0102() {
	fmt.Printf("concat abc def = %v", "abc"+"def")
}

func demo0103() {
	fmt.Print(`raw string 中没有转义，比如 \n, 但会删除 \r`)
}

func demo0104() {
	fmt.Println("coding = 世界")
	fmt.Println("coding = \xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("coding = \u4e16\u754c")
	fmt.Println("coding = \U00004e16\U0000754c")
}

func demo0105() {
	s := "Hello, 世界"
	fmt.Printf("string len = %v\n", len(s))
	fmt.Printf("rune count = %v\n", utf8.RuneCountInString(s))

	// 较笨拙的用法
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func demo0106() {
	s := "プログラム"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
}

func demo0107() {
	var buf bytes.Buffer
	buf.WriteByte('[')
	buf.WriteString("abc")
	buf.WriteByte(']')
	fmt.Fprintf(&buf, "%d", 5)
	fmt.Printf("buf.String() = %v\n", buf.String())
}
