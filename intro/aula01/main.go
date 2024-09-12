package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven > 0 {
		if actualMinutesInOven <= OvenTime {
			return OvenTime - actualMinutesInOven
		} else {
			return actualMinutesInOven - OvenTime
		}
	}
	return 0
}

func main() {
	fmt.Println(RemainingOvenTime(30))
}
