package main

import (	
	
	"lab365/calculadora"
	"lab365/impressora"
)

//letra maiúscula - função pública
//letra minúscula - função privada
func main(){
 
	res := calculadora.Soma(10,10)

	impressora.Imprimir(res)

	impressora.ImprimirString("oi", "ola","bom dia")

}