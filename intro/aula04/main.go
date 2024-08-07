package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://google.com.br")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Header)

	//
	ress, errs := soma(7, 3)

	if errs != nil {
		log.Fatal(errs.Error())
	}
	fmt.Println(ress)

	//
	ress2, _ := soma(10, 10)

	fmt.Println(ress2)

}

func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}
