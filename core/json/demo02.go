package main

import (
	"encoding/json"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}

// json 的 encoding 与 decoding
func demo0201() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	// 从 struct 到 byte[]
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))

	// 从 byte[] 到 struct
	var n Message
	if err = json.Unmarshal(b, &n); err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}

func demo0202() {
}
