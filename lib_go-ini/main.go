package main

import (
	"fmt"
	"lib_go-ini/config"
)

const (
	iniSource = "my.ini"
)

func main() {
	config := config.LoadConfigFromIni(iniSource)
	fmt.Printf("config = %#v\n", config)
}
