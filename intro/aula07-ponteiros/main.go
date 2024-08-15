package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "Kart"
	fmt.Println(c.Name, "andou")
}

func (c Carro) parou() {
	c.Name = "Chevete"
	fmt.Println(c.Name, "parou")
}

func main() {

	// example 1
	variavel := 10
	abc(&variavel)
	fmt.Println(variavel)

	// example 2
	carro := Carro{
		Name: "Ka",
	}

	carro.andou()
	fmt.Println(carro.Name)

	carro.parou()
	fmt.Println(carro.Name)

}

func abc(a *int) {
	*a = *a + 200
}
