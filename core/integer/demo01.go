package main

import (
	"fmt"
)

func demo0101() {
	// ^ 二元运算时，异或
	fmt.Printf("0011 ^ 0101 = %04b\n", 0b0011^0b0101)

	// 特别的规则，来自 golang spec
	// ^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
	//                                      and  m = -1 for signed x
	// all bits set to 1 是 11111111 11111111 11111111 11111111
	// -1 二进制为           11111111 11111111 11111111 11111111
	// 所以对 unsigned 或 signed 来说，规则没什么不同，只是看法上的区别

	// ^ 一元运算 无符号数
	var a byte = 0x0f
	fmt.Printf(" byte = %08b\n", a)
	fmt.Printf("^byte = %08b\n", ^a)

	// ^ 一元运算 有符号数
	var b int = 0x0f
	fmt.Printf(" 0x0f = %08b, %v\n", b, b)
	// fmt 输出时，作了转化，实际-16二进制为 11111111 11111111 11111111 11110000，^ 位取反，无论有无符号，操作是统一的
	fmt.Printf("^0x0f = %08b, %v\n", ^b, ^b)

	// 第二个操作数位1，结果为0，否则结果为第一个操作数位
	fmt.Printf("0b0011 &^ 0b0101 = %04b", 0b0011&^0b0101)
}
