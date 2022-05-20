package main

import "fmt"

type dailyMission struct{}

func NewDailyMission() *dailyMission {
	d := &dailyMission{}

	// 在初始化dailyMission的时候订阅'playLevelUp'主题
	// 即当有生产者向该主题publish内容时，该方法就会被调用
	_ = bus.Subscribe("playerLevelUp", d.OnPlayLevelUp)
	return d
}

func (d *dailyMission) OnPlayLevelUp(oldLevel, newLevel int) {
	fmt.Printf("daily mission oldLevel: %d, newLevel: %d\n", oldLevel, newLevel)
}
