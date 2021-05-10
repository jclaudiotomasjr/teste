package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome  string
	idade int
	email string
	sexo  string
}

func novaPessoa(nome string) *pessoa {
	p := pessoa{nome: nome}
	p.idade = 37
	p.email = strings.ToLower(p.nome + "@gmail.com")
	return &p
}

func main() {
	fmt.Println(pessoa{"Vitor", 11, "vitor@gmail.com", "M"})
	fmt.Println(pessoa{nome: "Yan", idade: 11, email: "yan@gmail.com", sexo: "M"})
	fmt.Println(pessoa{nome: "Jack"})
	fmt.Println(&pessoa{nome: "Jujuba", idade: 1})
	fmt.Println(novaPessoa("Junior"))

	j := pessoa{nome: "Stefania", idade: 39, email: "stefania@gmail.com", sexo: "F"}
	fmt.Println(j)

	jr := &j
	fmt.Println(jr.idade)

	jr.idade = 32
	fmt.Println(jr.idade)

}
