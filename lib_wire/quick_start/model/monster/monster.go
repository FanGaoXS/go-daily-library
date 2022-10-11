package monster

import "lib_wire/quick_start/environment"

type Monster struct {
	Name string
}

func New(env *environment.Env) *Monster {
	return &Monster{
		Name: env.Monster,
	}
}
