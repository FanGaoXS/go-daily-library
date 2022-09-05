package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	duration time.Duration
)

func main() {
	flag.Parse()

	fmt.Println("duration :", duration.String())
}

func init() {
	flag.DurationVar(&duration, "duration", time.Second, "duration of time")
}
