package main

import (
	"crypto/rand"
	"math/big"
)

func main() {
	for i := 0; i < 500; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(5000))
		println(n.Int64())
	}
}
