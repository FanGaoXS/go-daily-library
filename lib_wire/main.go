package main

func main() {
	args1 := monsterArgs{name: "mmm"}
	args2 := playerArgs{name: "mmm"}
	
	eA := InitEndingA(args1, args2)
	eA.ending()

	eB := InitEndingB(args1, args2)
	eB.ending()
}
