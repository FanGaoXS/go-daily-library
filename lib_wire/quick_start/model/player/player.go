package player

import "lib_wire/quick_start/environment"

type Player struct {
	Name string
}

func New(env *environment.Env) *Player {
	return &Player{
		Name: env.Player,
	}
}
