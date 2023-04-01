package main

import "fmt"

type Pessoa struct {
	Nome   string
	Cidade string
	Email  string
}

func main() {

	mapStruct := map[string]Pessoa{
		"044565904-16": Pessoa{
			Nome:   "Patrícia",
			Cidade: "São Paulo",
			Email:  "teste@teste.com",
		},
		"111111111-00": {
			Nome:   "Maria",
			Cidade: "São Paulo",
			Email:  "teste@teste.com",
		},
		"222222222-00": {
			Nome:   "Joaquina",
			Cidade: "São Paulo",
			Email:  "teste@teste.com",
		},
	}
	fmt.Println(mapStruct["044565904-16"])

	mapaEstados := map[string]string{
		"SP": "terra da garoa",
		"MG": "Terra do pao de queijo",
		"RJ": "Continua lindo",
		"BH": "",
	}

	valorChave, existeChave := mapaEstados["BH"]
		if existeChave {
			fmt.Println(valorChave)
		}

	for chave, valor := range mapaEstados {
		fmt.Println(chave, valor)
	}

	urlMap := map[string]string{
		"google":  "http://google.com",
		"youtube": "http://youtube.com",
		"globo":   "http://globo.com",
		"yahoo":   "http://yahoo.com",
		"bing":    "http://bing.com",
	}

	var input string
	
	fmt.Print("Digite a chave desejada: ")
	fmt.Scan(&input)

	url, ok := urlMap[input]

	if ok {
		fmt.Println(url)
	}else{
		fmt.Println("Chave não encontrada")
	}

}