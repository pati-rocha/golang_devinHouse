package main


import (
	"fmt"
	"math"
)

func main() {

	//tipos de dados

	//inteiros

	var inteiro int = 3
	var inteiro8 int8 = 23
	var inteiro16 int16 = 5666
	var inteiro32 int32 = 1234567891
	var inteiro64 int64 = 1234567891234567891
	

	fmt.Println(inteiro, inteiro8,inteiro16,inteiro32,inteiro64)

	//floats - pontos flutuantes, numeros com casas decimais

		var floatNumber float32 = math.Pi
		var floatNumber2 float64 = math.Pi
	
		fmt.Println(floatNumber, floatNumber2 )

	//booleanos

	var verdadeiro bool = true
	verdadeiroDois := true
	falso := false
	

	fmt.Println(verdadeiro, verdadeiroDois, falso)

	//strings - cadeia de caracter

	var saudacoes string = "Bem vindo ao Golang!"

	fmt.Println(saudacoes)

	//Runa ou Rune, guarda dados do tipo unicode. emojis
	var r rune = '😀'

	runaTexto := "Olá, 世界!" //Olá mundo em chinês

	fmt.Println(r, runaTexto)
	
}
