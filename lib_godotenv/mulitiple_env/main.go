package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".active-env"); err != nil {
		log.Fatalln(err)
	}

	// 根据不同环境去加载不同的配置文件到环境变量中
	activeEnv := os.Getenv("active")
	if activeEnv == "dev" || activeEnv == "" {
		if err := godotenv.Load(".env.dev"); err != nil {
			log.Fatalln(err)
		}
	} else if activeEnv == "test" {
		if err := godotenv.Load(".env.test"); err != nil {
			log.Fatalln(err)
		}
	} else if activeEnv == "prod" {
		if err := godotenv.Load(".env.prod"); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("app_name:", os.Getenv("app_name"))
}
