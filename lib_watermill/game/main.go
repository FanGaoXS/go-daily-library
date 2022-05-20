package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {

	manager := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	player := NewPlayer(manager)
	d := NewDailyMission(manager) // 消费者dailyMission：订阅主题'playerLevelUp'
	go d.OnPlayLevelUp()          // 消费者读取队列中的内容
	a := NewAchievement(manager)  // 消费者achievement：订阅主题'playerLevelUp'
	go a.OnPlayLevelUp()

	player.LevelUp() // 生产者：生产产品
	player.LevelUp()
	player.LevelUp()
}
