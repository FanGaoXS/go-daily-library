package main

import (
	"context"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

func main() {
	// 消息队列管理器
	manager := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	// 利用管理器订阅topic这一队列
	exampleQueue, err := manager.Subscribe(context.Background(), "example")
	if err != nil {
		panic(err)
	}

	go pullMsg(exampleQueue)

	pushMsg(manager)
}

// 将消息放入队列
func pushMsg(publisher message.Publisher) {
	for {
		// 消息内容，内容本身是[]byte类型（方便序列化对象）
		msg := message.NewMessage(watermill.NewUUID(), []byte("hello, world!"))
		// 向topic为'example'的队列放入msg
		if err := publisher.Publish("example", msg); err != nil {
			panic(err)
		}

		// 时间间隔1s发送一次
		time.Sleep(time.Second * 1)
	}
}

// 从队列中取出消息，该队列是channel类型
func pullMsg(messages <-chan *message.Message) {
	for msg := range messages {
		// 打印队列中的消息内容
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		// 确认收到消息后需要调用Ack()确认消息已经收到，否则会重发
		msg.Ack()
	}
}
