package main

import (
	"fmt"
	"log"
	"time"

	"github.com/araddon/dateparse"
)

func main() {
	// UTC
	dataStr := "3/1/2014 08:00:00"
	// mm/dd/yyyy
	t, err := dateparse.ParseAny(dataStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	location, _ := time.LoadLocation("America/Chicago")
	t1, _ := dateparse.ParseIn(dataStr, location)
	fmt.Println(t1.Local().Format("2006-01-02 15:04:05"))

	t2, _ := dateparse.ParseIn(dataStr, time.Local)
	fmt.Println(t2.Local().Format("2006-01-02 15:04:05"))
}
