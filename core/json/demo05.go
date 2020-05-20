package main

import (
	"encoding/json"
	"log"
	"os"
)

// 流处理
func demo0501() {
	// 标准输入
	dec := json.NewDecoder(os.Stdin)
	// 标准输出
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		// 从标准输入 decode 到 v
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
		// 可直接向控制台输出
		enc.Encode("{'console':'console'}")
	}

}
