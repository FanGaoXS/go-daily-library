package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// 只从 `.env` 文件中读，但不写入环境变量
	envMap, err := godotenv.Read(".env.yaml")
	if err != nil {
		log.Fatalln("read env err")
	}

	fmt.Println("active_env =", envMap["active_env"])
	fmt.Println("app.name =", envMap["app.name"])
	fmt.Println("app.creator =", envMap["app.creator"])
	fmt.Println("app.version =", envMap["app.version"])
}
