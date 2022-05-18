package main

import (
	"bytes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	buf := &bytes.Buffer{}
	buf.WriteString("name = awesome web")
	buf.WriteByte('\n')
	buf.WriteString("version = 0.0.1")

	env, err := godotenv.Parse(buf)
	if err != nil {
		log.Fatal(err)
	}

	// 写入到`./.env`路径中
	err = godotenv.Write(env, "./.env2")
	if err != nil {
		log.Fatal(err)
	}
}
