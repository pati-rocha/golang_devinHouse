package main

import "fmt"

func main(){

	//operadores aritméticos

	n1 := 10
	n2 := 2

	a := "oi"

	b := "tudo bem?"

	//op soma
	soma := n1 + n2
	ss := a + b 

	fmt.Println(soma)
	fmt.Println(ss)

	//subtração
	sub := n1 - n2

	fmt.Println(sub)

	//multiplicação

	mult := n1 * n2

	fmt.Println(mult)

	//divisão
	div := n1/n2

	fmt.Println(div)

	//módulo
	mod := n1 % n2

	fmt.Println(mod)

	//incremento
	n1++ //op unário

	fmt.Println(n1)

	//decremento
	n1--  //op unário

	fmt.Println(n1)

	//atribuições
	n1 = n1 + 1 //3 formas diferentes com resultados iguais
	n1 += 1
	n1++

	n1 -= 1
	n1 *= 1
	n1 /= 1
	n1 %= 2

}