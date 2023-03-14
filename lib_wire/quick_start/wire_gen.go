// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"lib_wire/quick_start/environment"
	"lib_wire/quick_start/model/mission"
	"lib_wire/quick_start/model/monster"
	"lib_wire/quick_start/model/player"
)

// Injectors from wire.go:

func InitMission(env *environment.Env) *mission.Mission {
	playerPlayer := player.New(env)
	monsterMonster := monster.New(env)
	missionMission := mission.New(playerPlayer, monsterMonster)
	return missionMission
}