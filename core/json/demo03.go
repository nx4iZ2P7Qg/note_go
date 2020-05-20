package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

// interface{} 通用类型
func demo0301() {
	var i interface{}
	i = "a string"
	i = 2011
	i = 2.777
	// type assertion
	r := i.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)
}

// 未知类型的处理
func demo0302() {
	var i interface{}
	i = "a string"
	//i = 5
	//i = 2.5
	// 获取 i 的实际类型
	switch v := i.(type) {
	case int:
		fmt.Println("twice i is", v*2)
	case float64:
		fmt.Println("the reciprocal of i is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("i swapped by halves is", v[h:]+v[:h])
	default:
		// i isn't one of the types above
	}
}

// 处理结构未知的 json
func demo0303() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Fatal("unmarshal b failed")
	}
	// map[string]interface{} 对应 json
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		// unmarshal() 方法得到的数据类型为 bool,float64,string,nil
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		// []interface{} 对应 array
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
