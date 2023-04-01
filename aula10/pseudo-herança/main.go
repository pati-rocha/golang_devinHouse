package main

import "fmt"

//Como acontece a pseudo-herança?
//struct embedding aninhamento
type Pessoa struct {
	//atributos
	Nome string 
	Cidade string
}

type Estudante struct {
	Pessoa
	Matricula int16
}

type Professor struct {
	Pessoa
	Salario float64
}

func main() {
	
	e1 := Estudante{Pessoa{"Maria","Patos"},1234}

	fmt.Println(e1.Nome)

	p1 := Professor{Pessoa{"Joao","Sousa"}, 99999.99}
	fmt.Println(p1.Salario)
}