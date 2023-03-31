package main


import (
	"fmt"
	"strconv"
	
)

func main() {


	//conversão entre tipos de dados

	var x int32 = 10
	var y int64 = int64(x) //chama o método e coloca a variável pra fazer o casting
	fmt.Println(y)

	
	var a float32 = 3.14 
	var b float64 = float64(a)
	fmt.Println(b) //3.140000104904175

	var c float32 = 2.5 
	//int16(c)
	var d int16 = int16(c)

	fmt.Println(d) //2

	//string para inteiro
	strQualquer := "777"

	numQualquer, err := strconv.Atoi(strQualquer)

	//controle de exceções
	//tratamento de erro
	if err != nil {
		fmt.Println("Erro na conversão: ", err)
		return
	}

	fmt.Println("Número inteiro: ", numQualquer) //777

	//inteiro para string
	var inteiro int = 101
	var texto string
	texto = strconv.Itoa(inteiro)

	fmt.Println("Conversão string: ",texto)


}
