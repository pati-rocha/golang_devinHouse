package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		//slice -> array dinâmico
		addArray := []int{0, 12, 34}
		//append -> adiciona ao final do array
		addArray = append(addArray, 10, 20, 30)

		fmt.Println(addArray)

		meuArray := [10]int{5, 3, 8, 6, 7, 5, 3}

		for i := 0; i < len(meuArray); i++ {
			//fmt.Print(meuArray[i])
		}

		/*
			var idade float64

			idade = 10.55
			for y, yy := range meuArray {
				fmt.Printf("\n %d : %d %f\n", y, yy, idade)
			}*/
	/*num := 0
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := []int{}

	//num := 9

	for i := 0; i <= len(arr); i++ {

		result = append(result, num*i)

		fmt.Printf("%d x  %d = %d \n", num, i, result[i])
	}*/

	carro := ""
	fmt.Print("Digite o carro: ")
	fmt.Scan(&carro)

	cars := []string{"mercedez","volvo","ferrari","bmw","ford",	"renault","vw"}

	
	for i, car := range cars{

		if car == strings.ToLower(carro){
			fmt.Printf("%s na posição %d. \n", car, i)
				
		} else {
		fmt.Println("Input inválido. Carro não encontrado!")
		}
		
	}
}
	

