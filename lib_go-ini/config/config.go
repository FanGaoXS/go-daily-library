package config

import (
	"github.com/go-ini/ini"
	"log"
)

type Mysql struct {
	ip       string
	port     string
	username string
	password string
	database string
}

type Redis struct {
	ip   string
	port string
}

type Config struct {
	appName     string
	logLevel    string
	mysqlConfig Mysql
	redisConfig Redis
}

func LoadConfigFromIni(source string) *Config {
	config, err := ini.Load("my.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	return &Config{
		appName:  config.Section("").Key("app_name").String(),
		logLevel: config.Section("").Key("log_level").String(),
		mysqlConfig: Mysql{ // 读ini文件的`[mysql]`板块
			ip:       config.Section("mysql").Key("ip").String(),
			port:     config.Section("mysql").Key("port").String(),
			username: config.Section("mysql").Key("username").String(),
			password: config.Section("mysql").Key("password").String(),
			database: config.Section("mysql").Key("database").String(),
		},
		redisConfig: Redis{ // 读ini文件的`[redis]`板块
			ip:   config.Section("redis").Key("ip").String(),
			port: config.Section("redis").Key("port").String(),
		},
	}
}
