package main

import (
	_ "lib_wire/env"
)

func main() {
	mission := InitMission()
	mission.Status()
	mission.Start()
	mission.Status()
}
