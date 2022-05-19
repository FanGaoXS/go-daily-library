package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
)

type Option struct {
	ConfigFile string `short:"c" long:"config" description:"Name of config file."`
	EnvFile    string `short:"e" long:"env" description:"Name of env file."`
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)
	if err != nil {
		return nil, err
	}

	return &opt, nil
}

type EnvOption struct {
	AppName    string
	AppCreator string
	AppVersion string
}

func InitEnv(opt *Option) (*EnvOption, error) {
	envMap, err := godotenv.Read(opt.EnvFile)
	if err != nil {
		return nil, err
	}

	return &EnvOption{
		AppName:    envMap["app_name"],
		AppCreator: envMap["app_creator"],
		AppVersion: envMap["app_version"],
	}, nil
}

func PrintEnvOption(opt *EnvOption) {
	fmt.Printf("env_option = %#v\n", opt)
}

func main() {
	// 1, 创建容器
	container := dig.New()

	// 2, 将对象的构造函数provide
	err := container.Provide(InitOption)
	if err != nil {
		log.Fatalln("provide InitOptions err")
	}
	err = container.Provide(InitEnv)
	if err != nil {
		log.Fatalln("provide InitEnv err")
	}

	// 3, 将函数需要的依赖从容器中注入
	container.Invoke(PrintEnvOption)

	// tips：需要特别注意的是，即使是provide简单的对象，也不能直接provide对象的地址，而是利用使用函数
	// 返回对象，然后provide该函数。
	// e.g：错误示范：
	//		u = NewUser()
	//		container.Provide(u) <=> container.Provider(NewUser()) :
	//
	// 	    正确示范：
	//		func initUser() *User{return NewUser()}
	//		container.Provide(initUser)
}
