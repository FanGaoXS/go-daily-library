package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Player  string
	Monster string
}

// Load config from environment file, and set them to environment variable
func Load() *Env {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("read environment file failed")
	}

	return &Env{
		Player:  os.Getenv("player"),
		Monster: os.Getenv("monster"),
	}
}

// Read just read config from environment file
func Read() *Env {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln("read environment file failed")
	}

	return &Env{
		Player:  envMap["player"],
		Monster: envMap["monster"],
	}
}
