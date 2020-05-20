package main

import "fmt"

// 如果 struct 内所有 field 都是 comparable，则 struct 也是 comparable，从而可以使用 == != 判断
func demo0101() {
	type Point struct{ X, Y int }
	p := Point{1, 2}
	q := Point{2, 1}
	r := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)                   // "false"
	fmt.Println(q == r)                   // true
}
