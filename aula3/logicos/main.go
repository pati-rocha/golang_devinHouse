package main

import "fmt"

func main(){

	//op lógicos
	//podemos utilizar todos os tipo de dados

	s1 := "go"
	s2 := "lang"

	//conjunção && - AND

	c := (s1 == "go" && s2 == "lang") //as duas condições precisam ser verdadeiras 

	fmt.Println(c)

	//disjunção || - OR
	dis := (s1 =="go" || s2 == "go") //uma ou outra condição deve ser verdadeira

	fmt.Println(dis)

	//negação !! 
	neg := (s1 != s2)

	fmt.Println(neg)

	

}