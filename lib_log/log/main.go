package main

import "log"

type user struct {
	name string
	age  int
}

func main() {
	u := &user{
		name: "test",
		age:  18,
	}
	// log的标识
	// 默认是 log.LstdFlags <=> log.Ldate | log.Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// Print/Printf/Println
	log.Printf("%#v\n", u)
	// Panic/Panicf/Panicln
	//log.Panicln(u)
	//// Fatal/Fatalf/Fatalln
	//log.Fatalln(u)
}
