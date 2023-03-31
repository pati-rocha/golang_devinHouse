package main

import "fmt"

func main(){

	//op relacionais o resultado da expressão será um valor booleano

	n1 := 12
	n2 := 100

	//comparação de igualdade
	isEqual := n1 == n2  //true or false

	fmt.Println(isEqual)

	//comparação de diferença
	isDiff := n1 != n2

	fmt.Println(isDiff)

	//maior
	isGreater := n1 > n2

	fmt.Println(isGreater)
	
	//menor
	isLess := n1 < n2

	fmt.Println(isLess)

	//maior que
	isGreaterOrEqual := n1 >= n2
	fmt.Println(isGreaterOrEqual )

	//menor que
	isLessOrEqual := n1 <= n2
	fmt.Println(isLessOrEqual)

	//usando string
	s1 := "go"
	s2 := "lang"

	str := s1 > s2

	fmt.Println(str)

	//nil

	var k *int8
	var c *int8
	c = nil
	k = nil
	fmt.Println("nil:", k == c)


}