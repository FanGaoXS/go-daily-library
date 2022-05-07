package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigType("toml")
	tomlConfig := []byte(`
app_name = "test web"

log_level = "DEBUG"

[mysql]
ip = "127.0.0.1"
port = 3306
user = "dj"
password = 123456
database = "awesome"

[redis]
ip = "127.0.0.1"
port = 6381

[server]
protocols = ["http", "https", "wss"]
ports = [8080, 8081, 8082]
timeout = "3s"`)
	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	fmt.Println("redis port: ", viper.GetInt("redis.port"))
}
