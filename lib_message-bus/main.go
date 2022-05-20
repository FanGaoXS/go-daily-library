package main

import (
	"fmt"
	messagebus "github.com/vardius/message-bus"
	"sync"
)

func main() {
	queueSize := 10
	bus := messagebus.New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	// 订阅topic
	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("v")
	})

	// 订阅topic
	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println("v")
	})

	// 向topic发送消息
	bus.Publish("topic", true)
	wg.Wait()
}
