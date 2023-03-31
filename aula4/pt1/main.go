package main

import "fmt"

func main() {

	str := "golang"
	fmt.Println("Valor da variável str : ", str) 
	fmt.Println("Valor do endereço da variável str: ", &str) 


	//para criar um ponteiro declara uma variável
	var ptr *string = &str //passando o endereço de memória da variável str

	//imprime o endereço do ponteiro/memória
	fmt.Println("o valor do ponteiro : ", ptr) 
	//imprime o valor do endereço do ponteiro
	fmt.Println(*ptr) 
	//imprime o endereço de memória
	fmt.Println(&str) 

	//ponteiros inicializados mas não utulizados são nil, o valor default é nil
	var i *int
	fmt.Println("o valor do ponteiro [i]: ", i) 

	//podemos inferir o tipo do ponteiro com var
	var n = 888
	var p = &n

	fmt.Println("o valor do  [p]: ", *p) 

	//podemos declarar com syntax breve
	n2 := 777
	p2 := &n2

	fmt.Println("o valor do [p2]: ", *p2) 

	//alterar o valor do ponteiro
	*p2 = 999
	fmt.Println("o valor do [n2]: ", n2) 



}