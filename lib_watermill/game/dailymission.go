package main

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
)

// 消费者
type dailyMission struct {
	msgQueue <-chan *message.Message
}

func NewDailyMission(subscriber message.Subscriber) *dailyMission {
	// 在初始化dailyMission的时候订阅'playLevelUp'主题
	// 即当有生产者向该主题publish内容时，msgQueue会有新消息
	msgQueue, err := subscriber.Subscribe(context.Background(), "playerLevelUp")
	if err != nil {
		panic(err)
	}

	d := &dailyMission{msgQueue: msgQueue}
	return d
}

// OnPlayLevelUp 循环从queue中取出消息
func (d *dailyMission) OnPlayLevelUp() {
	for m := range d.msgQueue {
		var msg map[string]interface{}
		_ = json.Unmarshal(m.Payload, &msg)
		log.Printf("daily mission received message: %s, payload: %v", m.UUID, msg)
		m.Ack() // 确认收到消息
	}
}
