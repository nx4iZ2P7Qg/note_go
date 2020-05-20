package main

import (
	"fmt"
	"math"
)

// 优先使用 float64，因为 float32 在计算中很快会出现问题
func demo0101() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!

	var f1 float64 = 16777216 // 1 << 24
	fmt.Println(f1 == f1+1)   // "false"!
}

// 最大最小值
func demo0102() {
	fmt.Printf("max float32 = %v\n", math.MaxFloat32)
	fmt.Printf("max float64 = %v\n", math.MaxFloat64)
	// 最小值约为1.4e-45 and 4.9e-324
}

// 特殊的数
func demo0103() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)

	fmt.Printf("math.NaN() = %v", math.IsNaN(math.NaN()))
}
