package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	log.Println("rand int", rand.Intn(100))
}
