package player

import "os"

type Player struct {
	Name string
}

func New() *Player {
	player := os.Getenv("player")
	return &Player{Name: player}
}
