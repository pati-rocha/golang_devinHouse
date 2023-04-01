// Estruturas de controle
package main

import "fmt"

type Pessoa struct {
	Nome string
}
func main() {

	/*age := 18
	elegibilidade := false
	planoVip := true

	if age >= 18 && elegibilidade {
		fmt.Println("Usuário liberado!")
	} else if age >= 18 && !elegibilidade && planoVip {
		fmt.Println("Pode entrar!")
	}
	//fmt.Println("Precisa ter 18 anos!")
	fmt.Println("Fim do código...") */

	/*num := 3

	if num == 5 || num == 3 {
		fmt.Println("Sucesso!")
	} else {
		fmt.Println("Erro")
	}*/

	/*num := 6

	switch num {
		case 1: 
			fmt.Println("Opção 1")
		case 2: 
			fmt.Println("Opção 2")
		case 3: 
			fmt.Println("Opção 3")
		case 4: 
			
		case 5: 
			fmt.Println("Opção 4 ou 5")
		default: 
			fmt.Println("Opção inválida")
	}*/

	/*lista := [3]int{1,2,3}

	for _, i := range lista {
		fmt.Println("oi", i)
	}
	for i := range lista {
		fmt.Println("oi oi", lista[i])
	}

	for i := 0; i < len(lista); i++ {
		fmt.Println("oi de novo", i)
	}*/

	//array := [6]int{1,1,3,4,5,6} //finito
	//slice := make(map[string]int) //dinâmico
	
	//map é uma coleção de chave e valor
	mapEx := map[string]string{
		"chave1": "valor um",
		"chave2": "valor dois",
	} 
	mapEx["chave1"] = "oi"
	fmt.Println(mapEx["chave1"])


	mapExx := map[string]int{} 
	mapExx["chave1"] = 1
	mapExx["chave2"] = 10
	mapExx["chave3"] = 100
	

	teste := mapExx["chave3"]
	fmt.Println(teste)


	mapExemplo := map[string]interface{}{}

	mapExemplo["chave2"] = 2
	mapExemplo["chave3"] = "três"
	mapExemplo["chave4"] = true

	fmt.Println(mapExemplo["chave2"])
	fmt.Println(mapExemplo["chave3"])
	fmt.Println(mapExemplo["chave4"])

	mapStruct := map [string]Pessoa{
		"044565904-16": Pessoa{
			Nome: "Patrícia",
		},
	}
	fmt.Println(mapStruct["044565904-16"])





}