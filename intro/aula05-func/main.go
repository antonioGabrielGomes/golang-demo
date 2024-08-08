package main

import (
	"fmt"
)

func main() {
	r, m := soma(1, 2)
	fmt.Println(r, m)

	r2 := soma2(10, 30)
	fmt.Println(r2)

	r3 := soma3(10, 20, 30)
	fmt.Println(r3)

	r4 := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}

	fmt.Println(r4(10, 20, 30)())
}

func soma(x int, y int) (int, string) {
	return x + y, "somou"
}

func soma2(x int, y int) (result int) {
	result = x + y
	return
}

func soma3(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}

	return resultado
}
