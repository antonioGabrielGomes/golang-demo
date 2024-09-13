package main

import (
	"fmt"
)

// arrays
func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	// strconv.QuoteToASCII
	fmt.Println(a[0][0])
	fmt.Println(string(a[0][0]))

	//fmt.Println(a[0], a[1])
	//fmt.Println(a)

	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)
}

// loop
func loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
