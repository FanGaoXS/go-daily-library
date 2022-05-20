package main

import (
	messagebus "github.com/vardius/message-bus"
	"time"
)

const queueSize = 10

var bus = messagebus.New(queueSize)

func main() {
	p := NewPlayer()
	NewAchievement()  // 消费者：订阅'playerLevelUp'的消息
	NewDailyMission() // 消费者：订阅'playerLevelUp'的消息

	p.LevelUp() // 生产者
	p.LevelUp()
	p.LevelUp()

	time.Sleep(time.Second * 1)
}
