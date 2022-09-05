package main

import (
	"github.com/google/wire"
	"lib_wire/mission"
	"lib_wire/monster"
	"lib_wire/player"
)

func InitMission() *mission.Mission {
	wire.Build(monster.New, player.New, mission.New)
	return &mission.Mission{}
}
