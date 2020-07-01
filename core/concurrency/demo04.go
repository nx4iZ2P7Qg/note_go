package main

import "fmt"

// 只写单向通道
func demo0401() {
	c := make(chan int)

	// 隐式，通道转单向通道，单向无法转通道
	go test(c)

	for i := range c {
		fmt.Printf("%v\n", i)
	}
}

func test(c chan<- int) {
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
