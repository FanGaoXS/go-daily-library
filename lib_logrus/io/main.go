package main

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

func main() {
	io1 := &bytes.Buffer{}
	io2 := os.Stderr
	io3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(io1, io2, io3))
	logrus.Info("info msg")
}
