package impressora

import "fmt"


func ImprimirString(mensagem ...string) {
	fmt.Println()
}


func Imprimir(msg interface{}) {
	fmt.Println(msg)
}