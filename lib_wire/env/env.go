package env

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// load from '.env' then write into path
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("load env err")
	}
}
