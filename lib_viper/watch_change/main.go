package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	viper.SetConfigName("my")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	// 监听配置文件的修改
	viper.WatchConfig()

	// 配置文件修改的回调函数，即监听到配置文件发生修改，则会执行该函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})

	time.Sleep(time.Second * 10)
}
