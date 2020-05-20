package main

import (
	"fmt"
	"log"
	"note_go/common"
	"sort"
)

// 基本操作
func demo0101() {
	m1 := make(map[string]int)
	m2 := map[string]int{}
	m1["m1"] = 1
	m2["m2"] = 2

	val, ok := m1["m1"]
	common.CheckError(ok)
	log.Println(val)

	delete(m2, "m2")

	val, ok = m2["m2"]
	common.CheckError(ok)
}

// 有序取出 map 元素
func demo0102() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"bob":     32,
	}
	// 取 key
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	// key 排序
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

// 将 slice 作为 key
func demo0103() {
	// 第一步，定义函数 helper function k 将 [] 转成 string，同时满足 k(x) == k(y) 当且仅当 x == y
	// 第二步，定义 map[k()]interface{}
	var s []string
	Add(s)
	Add(s)
	log.Print(Count(s))
}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

// 使用这个方法，任何非 comparable 都可以作 key
// 想定义 非 == 的 equality 时可以使用
// k() 不一定是 string，其他 comparable 都可以
