package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// 将网页返回 json 以字符串、 map 形式打印出来
func demo0101() {
	resp, err := http.Get("http://httpbin.org/json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// 从 io.Reader 到 byte[]
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 从 byte[] 到 string
		bodyString := string(bodyBytes)
		log.Println("bodyString", bodyString)

		// 从 byte[] 到 map
		var jsonMap map[string]interface{}
		err = json.Unmarshal(bodyBytes, &jsonMap)
		if err != nil {
			log.Fatal("ERROR:", err)
		}
		log.Println("jsonMap", jsonMap)
	} else {
		log.Fatal("response status error")
	}
}
