package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 1 godotenv加载文件，并且set到环境变量中
	// 2 通过os.Getenv(key)根据key读取环境变量value

	err := godotenv.Load(".env", ".env2")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("app_name:", os.Getenv("app_name"))
	fmt.Println("app_version:", os.Getenv("app_version"))
}
