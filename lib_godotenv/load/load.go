package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// 从 `.env` 文件中读取，并且写入系统环境变量
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("load env err")
	}

	switch os.Getenv("active_env") {
	case "dev":
		err = godotenv.Load(".env-dev")
		if err != nil {
			log.Fatalln("load env-dev err")
		}
	case "test":
		err = godotenv.Load(".env-test")
		if err != nil {
			log.Fatalln("load env-test err")
		}
	case "prod":
		err = godotenv.Load(".env-prod")
		if err != nil {
			log.Fatalln("load env-prod err")
		}
	}

	// 利用os.Getenv(key)从环境变量中读取
	name := os.Getenv("name")
	appName := os.Getenv("app_name")
	appCreator := os.Getenv("app_creator")
	appVersion := os.Getenv("app_version")

	fmt.Printf("name = %#v\n", name)
	fmt.Printf("app_name = %#v\n", appName)
	fmt.Printf("app_creator = %#v\n", appCreator)
	fmt.Printf("app_version = %#v\n", appVersion)
}
