//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 将可能重复使用的依赖合并
var monsterPlayerSet = wire.NewSet(newPlayer, newMonster)

func InitEndingA(args1 monsterArgs, args2 playerArgs) *EndingA {
	wire.Build(monsterPlayerSet, NewEndingA)
	return &EndingA{}
}

func InitEndingB(args1 monsterArgs, args2 playerArgs) *EndingB {
	wire.Build(monsterPlayerSet, NewEndingB)
	return &EndingB{}
}
