package main

import "fmt"

func main() {

	//declarando canal, qdo se coloca tamanho no canal como buffer
	//buffer - reter msgs que ainda precisam de direcionamento, leitura
	ch := make(chan string, 5)

	//enviando msg para o canal
	ch <- "msg1"
	fmt.Println(<-ch)
	ch <- "msg2"
	ch <- "msg3"
	ch <- "msg4"
	ch <- "msg5"

	//channel implenta lógica FIFO = first in first out
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
}