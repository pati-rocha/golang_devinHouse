package main

import (
	"fmt"
	"time"
)

func main(){
	
	//blocking channel - bloqueio de canal
	buffer := make(chan int, 2)

	//envio de dado 1
	buffer <- 1
	buffer <- 4
	
	//rotina de sincronização entre sender e receiver
	//esta rotina monitora a presença de novos elementos no buffer(channel)
	go func(){ 
		for a := range buffer {
			fmt.Println(a)
		}
	}()

	//envio de dado 2
	buffer <- 2
	buffer <- 3
	time.Sleep(time.Second * 2)
}