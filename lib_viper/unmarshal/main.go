package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	AppName  string
	LogLevel string

	Mysql MysqlEnv
	Redis RedisEnv
}

type MysqlEnv struct {
	Ip       string
	Port     int
	User     string
	Password string
	database string
}

type RedisEnv struct {
	Ip   string
	Port int
}

func main() {
	viper.SetConfigName("my")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	var env Env
	viper.Unmarshal(&env)
	fmt.Println("env =", env)
}
