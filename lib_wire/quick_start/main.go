package main

import (
	"lib_wire/quick_start/environment"
)

func main() {
	env := environment.Read()
	mission := InitMission(env)
	mission.Start()
}
