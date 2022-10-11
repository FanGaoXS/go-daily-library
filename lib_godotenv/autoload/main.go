package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println("app_name:", os.Getenv("app_name"))
	fmt.Println("app_version:", os.Getenv("app_version"))
}
