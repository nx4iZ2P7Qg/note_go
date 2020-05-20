package main

import "fmt"

// Printf 中参数的复用，# 的使用
func demo0101() {
	fmt.Printf("%d, %[1]o, %#[1]x", 16)
}
