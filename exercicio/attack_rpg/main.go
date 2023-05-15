package main

import "fmt"

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	res := CanFastAttack(knightIsAwake)
	res = CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake)

	fmt.Println(res)
	fmt.Println(res)
}

func CanFastAttack(b bool) bool {
	if b == true {
		return false
	} else {
		return true
	}
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake == false && archerIsAwake == true && prisonerIsAwake == false {
		return true
	} else {
		return false
	}
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake == false && prisonerIsAwake == true {
		return true
	} else {
		return false
	}
}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if knightIsAwake == false && archerIsAwake == true && prisonerIsAwake == false && prisonerIsAwake == false {
		return false
	} else {
		return true
	}
}
