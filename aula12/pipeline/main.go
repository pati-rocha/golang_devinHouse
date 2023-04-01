package main

import "fmt"

func main() {

	//pipelines são formas ou frameworks para organizarmos o fluxo de trabalho

	//input de dados
	n := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//stage 1 : transformar input (slice) em channel
	datach := sliceToChannel(n)

	//stage 2 : processing data (elevar ao quadrado)
	productCh := sq(datach)

	//stage 3 : consume data
	for n := range productCh {
		fmt.Println(n)
	}
}

func sliceToChannel(n []int) <-chan int{

	out := make(chan int) //unbuffered channel

	//rotina que escreve no primeiro canal
	go func (){
		for _, n := range n {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(ch <-chan int) <-chan int {

	out := make(chan int) //unbuffered channel

	//rotina que escreve no primeiro canal
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()
	return out

}