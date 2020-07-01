package main

import (
	"fmt"
	"time"
)

func demo0101() {
	// 每100ms往通道中发1个信号
	tick := time.Tick(100 * time.Millisecond)
	// 500ms后往通道中发1个信号
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
