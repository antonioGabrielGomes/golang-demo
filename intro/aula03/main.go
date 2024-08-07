package main

import (
	"aula03/op"
	"fmt"
)

func main() {
	a := 10
	b := "world"
	c := 3.144
	d := false
	e := `uouuuu este`

	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T \n", b, b)
	fmt.Printf("%v %T \n", c, c)
	fmt.Printf("%v %T \n", d, d)
	fmt.Printf("%v %T \n", e, e)

	resultado := op.Soma(1, 1)
	fmt.Printf("%v %T \n", resultado, resultado)

	resultadoSub := op.Sub(2)
	fmt.Printf("%v %T \n", resultadoSub, resultadoSub)
}
