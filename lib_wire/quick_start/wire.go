//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"lib_wire/quick_start/environment"
	"lib_wire/quick_start/model/mission"
	"lib_wire/quick_start/model/monster"
	"lib_wire/quick_start/model/player"
)

func InitMission(env *environment.Env) *mission.Mission {
	wire.Build(
		player.New,
		monster.New,
		mission.New,
	)
	return &mission.Mission{}
}
