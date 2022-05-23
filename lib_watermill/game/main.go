package main

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {

	logger := watermill.NewStdLogger(false, false)

	// channel
	goChannel := gochannel.NewGoChannel(gochannel.Config{}, logger)

	player := NewPlayer(goChannel)
	d := NewDailyMission(goChannel) // 消费者dailyMission：订阅主题'playerLevelUp'
	go d.OnPlayLevelUp()            // 消费者读取队列中的内容
	a := NewAchievement(goChannel)  // 消费者achievement：订阅主题'playerLevelUp'
	go a.OnPlayLevelUp()

	for {
		var number int
		n, _ := fmt.Scanln(&number)
		if n > 0 {
			switch number {
			case 1:
				player.LevelUp() // 生产者：生产产品
			case 0:
				return
			}
		}
	}
}
