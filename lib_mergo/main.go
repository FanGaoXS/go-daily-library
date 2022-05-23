package main

import (
	"fmt"
	"github.com/imdario/mergo"
	"log"
)

type Config struct {
	Address string
	Port    int
}

func main() {
	c1 := &Config{
		Address: "localhost",
		Port:    8080,
	}

	// 将c1合并到c2，struct
	var c2 Config
	if err := mergo.Merge(&c2, c1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("merge to struct = %#v\n", c2)

	// 将c1合并到c3，map
	c3 := make(map[string]interface{})
	if err := mergo.Map(&c3, c1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("merge to map = %#v\n", c3)

	// 在merge的时候，可以使用mergo.WithOverride参数，来执行merge时候的覆盖属性
	// mergo.WithAppendSlice参数，执行merge时，不覆盖slice属性，而是append
	// mergo.WithTypeCheck参数，切片的类型不一致时会出错
}
