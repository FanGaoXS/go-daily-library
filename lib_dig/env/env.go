package env

import (
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

type UserEnv struct {
	Name string
	Age  int
}

type ShoeEnv struct {
	Brand string
}

type Option struct {
	User *UserEnv
	Shoe *ShoeEnv
}

func InitEnv() *Option {
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalln("read env err")
	}

	age, _ := strconv.Atoi(envMap["user_age"])

	return &Option{
		User: &UserEnv{
			Name: envMap["user_name"],
			Age:  age,
		},
		Shoe: &ShoeEnv{
			Brand: envMap["shoe_brand"],
		},
	}
}
