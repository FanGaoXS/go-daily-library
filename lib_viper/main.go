package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// 文件名称（不需要带文件类型的后缀）
	viper.SetConfigName("my")
	// 文件类型
	viper.SetConfigType("toml")
	// 文件路径
	viper.AddConfigPath(".")
	// 读取文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	// Get(key), 根据键获取值
	fmt.Println("app_name =", viper.Get("app_name"))
	fmt.Println("log_level =", viper.Get("log_level"))

	// GetType(key), 返回指定类型的值，如果不存在则返回该类型的0值
	fmt.Println("mysql.ip =", viper.GetString("mysql.ip"))
	fmt.Println("mysql.port =", viper.Get("mysql.port"))
	fmt.Println("mysql.user =", viper.Get("mysql.user"))
	fmt.Println("mysql.password =", viper.Get("mysql.password"))
	fmt.Println("mysql.database =", viper.Get("mysql.database"))

	fmt.Println("redis.ip =", viper.Get("redis.ip"))
	fmt.Println("redis.port =", viper.Get("redis.port"))

	// IsSet(key), 判断某个键是否存在
	if viper.IsSet("app_name") {
		fmt.Println("app_name is set")
	} else {
		fmt.Println("app_name is not set")
	}

	// GetStringMap(key), 返回某个键下的所有键值对
	serverMap := viper.GetStringMap("server")
	protocols := serverMap["protocols"]
	fmt.Println("protocols =", protocols)
	ports := serverMap["ports"]
	fmt.Println("ports =", ports)
	timeout := viper.GetDuration("server.timeout")
	fmt.Println("timeout =", timeout)

	// AllSettings(), 返回所有的设置
	fmt.Println("all settings =", viper.AllSettings())
}
