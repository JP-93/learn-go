package main

import "fmt"

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	var petDogIsPresent = false

	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))

}

func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	if archerIsAwake == false && petDogIsPresent == true {
		return true
	} else if prisonerIsAwake == true && knightIsAwake == false && archerIsAwake == false {
		return true
	}

	return false
}
