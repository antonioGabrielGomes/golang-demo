package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Nome  string
	Email string
	CPF   int
}

func (c Client) Exibe() {
	fmt.Println("Exibindo cliente pelo metodo: ", c.Nome)
}

type ClientInternacional struct {
	Client
	Pais string `json:"pais"`
}

func main() {
	// cliente 1
	client := Client{
		Nome:  "pombocoelho",
		Email: "pombocoelho@pc.com",
		CPF:   12345678996,
	}
	fmt.Println(client)
	client.Exibe()

	// cliente 2
	client2 := Client{"mari", "m@m.com", 98778909889}
	fmt.Printf("Nome: %s. Email: %s. CPF: %d\n", client2.Nome, client2.Email, client2.CPF)
	client2.Exibe()

	// cliente 3
	client3 := ClientInternacional{
		Client: Client{
			Nome:  "pombocoelho",
			Email: "pombocoelho@pc.com",
			CPF:   12345678996,
		},
		Pais: "BR",
	}
	fmt.Printf("Nome: %s. Email: %s. Pais: %s\n", client3.Nome, client3.Email, client3.Pais)
	client3.Exibe()

	// cliente json
	clienteJson, err := json.Marshal(client3)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(clienteJson)) // clienteJson é do tipo []byte

	// cliente 4
	jsoncliente4 := `{"Nome":"pombocoelho","Email":"pombocoelho@pc.com","CPF":12345678996,"pais":"BR"}`
	cliente4 := ClientInternacional{}

	json.Unmarshal([]byte(jsoncliente4), &cliente4) // passagem de cliente4 por referencia
	fmt.Println(cliente4)

}
