package main

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
)

// 消费者
type achievement struct {
	msgQueue <-chan *message.Message
}

func NewAchievement(subscriber message.Subscriber) *achievement {
	msgQueue, err := subscriber.Subscribe(context.Background(), "playerLevelUp")
	if err != nil {
		panic(err)
	}

	a := &achievement{msgQueue: msgQueue}
	return a
}

func (a *achievement) OnPlayLevelUp() {
	for m := range a.msgQueue {
		var msg map[string]interface{}
		_ = json.Unmarshal(m.Payload, msg)
		log.Printf("achievement received message: %s, payload: %v", m.UUID, msg)
		m.Ack()
	}
}
