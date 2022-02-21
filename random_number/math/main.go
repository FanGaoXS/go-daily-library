package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 4; i++ {
		println(rand.Int63n(100))
	}
}
