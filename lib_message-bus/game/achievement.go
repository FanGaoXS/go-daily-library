package main

import "fmt"

type achievement struct{}

func NewAchievement() *achievement {
	a := &achievement{}

	// 在初始化achievement的时候订阅'playLevelUp'主题
	// 即当有生产者向该主题publish内容时，该方法就会被调用
	_ = bus.Subscribe("playerLevelUp", a.OnPlayLevelUp)
	return a
}

func (a *achievement) OnPlayLevelUp(oldLevel, newLevel int) {
	fmt.Printf("achievement oldLevel: %d, newLevel: %d\n", oldLevel, newLevel)
}
