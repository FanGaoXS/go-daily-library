package main

type Player struct {
	level int
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) LevelUp() {
	oldLevel := p.level
	p.level++
	newLevel := p.level

	// 向topic为'playerLevelUp'的队列中发送消息
	bus.Publish("playerLevelUp", oldLevel, newLevel)
}
