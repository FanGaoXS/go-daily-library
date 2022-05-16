package main

import (
	"fmt"
	"github.com/spf13/cast"
	"log"
)

func main() {
	// cast.ToType()
	fmt.Println(cast.ToString(true)) // true
	fmt.Println(cast.ToInt(true))    // 1
	fmt.Println(cast.ToBool(1))      // true

	// cast.ToTypeE()
	s, err := cast.ToStringE(true)
	if err != nil {
		log.Fatalln("conversion failed.")
	}
	fmt.Println(s) // true
	i, err := cast.ToInt64E(true)
	if err != nil {
		log.Fatalln("conversion failed.")
	}
	fmt.Println(i) // 1
	b, err := cast.ToBoolE(1)
	if err != nil {
		log.Fatalln("conversion failed.")
	}
	fmt.Println(b) // true
}
