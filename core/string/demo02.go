package main

import (
	"fmt"
	"io"
	"log"
	"note_go/common"
	"strconv"
	"strings"
)

// int string互转
func demo0201() {
	i, err := strconv.Atoi("-3")
	if err != nil {
		log.Fatal("error")
	}
	log.Println("i + 1 = ", i+1)
	s := strconv.Itoa(7)
	log.Println("s + a = ", s+"a")
}

// string 转其他
func demo0202() {
	b, err1 := strconv.ParseBool("true")
	common.CheckError(err1)
	log.Println(b)

	f, err2 := strconv.ParseFloat("3.1415", 64)
	common.CheckError(err2)
	log.Println(f)

	i, err3 := strconv.ParseInt("-42", 10, 64)
	common.CheckError(err3)
	log.Println(i)

	u, err4 := strconv.ParseUint("42", 10, 64)
	common.CheckError(err4)
	log.Println(u)
}

// 其他 转 string
func demo0203() {
	s1 := strconv.FormatBool(true)
	log.Println(s1)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	log.Println(s2)
	s3 := strconv.FormatInt(-42, 16)
	log.Println(s3)
	s4 := strconv.FormatUint(42, 16)
	log.Println(s4)
}

// string byte[] 转换
func demo0204() {
	s := "abc"
	// 一般来说需要复制数据到 []byte
	b := []byte(s)
	// 再次复制数据到 string
	s2 := string(b)
	fmt.Printf("s2 = %v\n", s2)
}

// 生成 string 对应的 Reader
func demo0205() {
	s := "ab"
	var r io.Reader
	r = strings.NewReader(s)
	fmt.Println(r)
}
