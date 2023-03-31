//denota/declara o pacote ao qual pertence o código.
package main //main.go -> entrypooint do programa. Por onde o programa começa a ser executado.

//importação de bibliotecas.
import (

	"fmt"
)

func main() {
	
	fmt.Println("Olá, qual o seu nome?")

	var name string
	//escutar prompt comando para input de dados do usuário.
	fmt.Scanln(&name) //&name -> endereço da memória onde a variável está alocada.

	fmt.Printf("Seja bem vindo(a), %s! ", name)

}