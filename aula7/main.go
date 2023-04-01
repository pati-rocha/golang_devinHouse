package main

import "fmt"

func main() {

	mensagem := " olá golang"

	//outra forma de declarar função, lembrar do escopo de função
	imprimirMensagem := func (mens1 string) (string , int, bool){
		valorRecebido := mens1 + mens1 + mens1
			//return mens1 + mens1 + mens1	
		return valorRecebido, 10, true	
	}
	resposta, numero, boleano := imprimirMensagem(mensagem )
	fmt.Println(resposta,  numero, boleano)
}

func imprimirMensagem( mens1 string) (string , int, bool){
	valorRecebido := mens1 + mens1 + mens1
	//return mens1 + mens1 + mens1
	return valorRecebido, 10, true
}

//outra forma declarar função, retorno nomeados
/*func imprimirMensagem( mens1 string) (msg string , valor int, ok bool){

	msg = mens1 + mens1 + mens1
	valor = 100
	ok = false

	//return mens1 + mens1 + mens1

	return 
}*/