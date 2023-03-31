package main

import "fmt"

//esta função recebe como parâmetro um ponteiro guardado em s (nome do ponteiro)
func AlteraValorComPonteiro(s *string) {

	//acontecem varias coisa...
	//no final altera-se o parametro recebido inicialmente
	fmt.Println("O endereco do ponteiro é: ", s)
	fmt.Println(*s)
	*s = "novo valor"
	fmt.Println(*s)

}
func AlteraValorSemPonteiro(s string) string {

	fmt.Println("O endereco da str da funcao : ", &s)
	fmt.Println(s)
	s = "novo valor"
	fmt.Println("novo valor da string: ", s)
	return s

}
func main() {

	str := "golang"

	AlteraValorComPonteiro(&str)

	//fmt.Println("o endreço da string é: ", &str)
	//fmt.Println(str)

	fmt.Println("o endreço da string é: ", &str)

	str = AlteraValorSemPonteiro(str)
	fmt.Println(str)

}