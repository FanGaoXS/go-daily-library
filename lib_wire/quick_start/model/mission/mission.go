package mission

import (
	"fmt"
	"lib_wire/quick_start/model/monster"
	"lib_wire/quick_start/model/player"
)

type Mission struct {
	player  *player.Player
	monster *monster.Monster
}

func New(p *player.Player, m *monster.Monster) *Mission {
	return &Mission{
		player:  p,
		monster: m,
	}
}

func (m Mission) Start() {
	fmt.Printf("%s killed the %s", m.player.Name, m.monster.Name)
}
