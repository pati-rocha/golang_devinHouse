package main

import "fmt"

//Implementar canais(channels)

func main() {

	//criar um canal sem buffer
	ch := make(chan string)

	//enviando msgs para o canal

	go func ()  {
		ch <- "Mensagem inicial"
	}()

	//recuperando msg do canal
	msg := <-ch
	fmt.Println(msg)

	//fechando canal
	close(ch)

}