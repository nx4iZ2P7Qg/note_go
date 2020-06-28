package myscript

import (
	"log"
	"time"
)

// 计算从2020-02-27开始到当天，戒烟天数
func Smoke() {
	start, err := time.ParseInLocation("2006-01-02", "2020-02-27", time.Local)
	if err != nil {
		log.Fatal("error", err)
	}
	duration := time.Now().Sub(start)

	log.Printf("duration = %v", duration.Hours()/24)
}
