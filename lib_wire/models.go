package main

import "fmt"

type monster struct {
	name string
}

type monsterArgs struct {
	name string
}

func newMonster(args monsterArgs) *monster {
	return &monster{name: args.name}
}

type player struct {
	name string
}

type playerArgs struct {
	name string
}

func newPlayer(args playerArgs) *player {
	return &player{name: args.name}
}

type mission struct {
	p *player
	m *monster
}

func newMission(p *player, m *monster) *mission {
	return &mission{
		p: p,
		m: m,
	}
}

func (m mission) fight() {
	fmt.Printf("%s(player) is fighting with %s(monster)\n", m.p.name, m.m.name)
}

type EndingA struct {
	m *mission
}

func NewEndingA(p *player, m *monster) *EndingA {
	return &EndingA{m: newMission(p, m)}
}

func (e EndingA) ending() {
	e.m.fight()
	fmt.Println("mission failed")
}

type EndingB struct {
	m *mission
}

func NewEndingB(p *player, m *monster) *EndingB {
	return &EndingB{m: newMission(p, m)}
}

func (e EndingB) ending() {
	e.m.fight()
	fmt.Println("mission success")
}
