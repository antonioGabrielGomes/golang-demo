package main

import "fmt"

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
	return actualMinutesInOven + PreparationTime(numberOfLayers)
}

func main() {
	// 0 40
	fmt.Printf("RemainingOvenTime: %d", RemainingOvenTime(40))
	fmt.Println()
	fmt.Printf("PreparationTime: %d", PreparationTime(0))
	fmt.Println()
	fmt.Printf("ElapsedTime: %d", ElapsedTime(0, 40))

	// 1 30
	fmt.Printf("RemainingOvenTime: %d", RemainingOvenTime(30))
	fmt.Println()
	fmt.Printf("PreparationTime: %d", PreparationTime(1))
	fmt.Println()
	fmt.Printf("ElapsedTime: %d", ElapsedTime(1, 30))
}
