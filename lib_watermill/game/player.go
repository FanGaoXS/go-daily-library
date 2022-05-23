package main

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"time"
)

// Player 生产者
type Player struct {
	level     int
	publisher message.Publisher
}

func NewPlayer(publisher message.Publisher) *Player {
	return &Player{
		level:     0,
		publisher: publisher,
	}
}

func (p *Player) LevelUp() {
	oldLevel := p.level
	p.level++
	newLevel := p.level

	bytes, _ := json.Marshal(map[string]interface{}{
		"old_level": oldLevel,
		"new_level": newLevel,
	})
	msg := message.NewMessage(watermill.NewUUID(), bytes)
	if err := p.publisher.Publish("playerLevelUp", msg); err != nil {
		panic(err)
	} // 向topic为'playerLevelUp'的队列中发送消息

	time.Sleep(time.Second * 1)
}
