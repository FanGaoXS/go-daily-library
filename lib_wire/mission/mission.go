package mission

import (
	"fmt"
	"lib_wire/monster"
	"lib_wire/player"
	"os"
)

type Mission struct {
	p           *player.Player
	m           *monster.Monster
	description string
	isFinished  bool
	isSuccess   bool
}

func New(p *player.Player, m *monster.Monster) *Mission {
	desc := os.Getenv("mission")
	return &Mission{
		p:           p,
		m:           m,
		description: desc,
		isFinished:  false,
		isSuccess:   false,
	}
}

func (m *Mission) Start() {
	if !m.isFinished {
		m.isSuccess = true
		m.isFinished = true
		fmt.Println("mission:", m.description, "start and success")
	} else {
		fmt.Println("mission:", m.description, "is finished")
	}
}

func (m Mission) Status() {
	fmt.Println("mission:", m.description)
	fmt.Println("is finished:", m.isFinished)
	fmt.Println("is success:", m.isSuccess)
}
