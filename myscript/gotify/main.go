package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"note_go/common"
	"os"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	c := cron.New()
	c.AddFunc("*/10 8-23 * * *", smart)
	wg.Add(1)
	c.Start()

	wg.Wait()
}

func smart() {
	b := strExistInFile("/var/log/messages", "connection status changed")
	if b {
		requestGotify("A61-EfaWFIMejMf", "hdd warning", "connection status changed", 10)
	}
}

func strExistInFile(file string, str string) bool {
	f, err := os.Open(file)
	common.CheckError(err)
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		b := strings.Contains(sc.Text(), str)
		if b {
			return true
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return false
}

func requestGotify(token string, title string, message string, priority int) {
	url := fmt.Sprintf("https://svrx.asuscomm.com:3313/message?token=%v", token)
	var data = make(map[string]interface{})
	data["title"] = title
	data["message"] = message
	data["priority"] = priority
	bytesData, err := json.Marshal(data)
	common.CheckError(err)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{
		Transport: tr,
	}
	res, err := c.Post(url, "application/json;charset=utf-8", bytes.NewBuffer(bytesData))
	common.CheckError(err)
	defer res.Body.Close()

	//content, err := ioutil.ReadAll(res.Body)
	//common.CheckError(err)

	//str := (*string)(unsafe.Pointer(&content))
	//fmt.Println(*str)
}
