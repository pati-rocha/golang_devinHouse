//qual é o motivo deste padrão? Qual dor ele resolve?
//serve para garantir o fechamento de uma rotina que não tem mais canal
//quando um canal se fecha e avisa suas rotinas dependentes que precisam ser encerradas

package main

import (
	"fmt"
	"time"
)

func main(){

	//canal que avisa o encerramento dos trabalhos ou a continuidade
	done := make(chan bool)

	//rotina para processamento de dados
	go processData(done)
	
	time.Sleep(time.Second * 1)

	//fechando o canal: todas as vezes que fechamos canais priduzimos mensagem
	close(done)

	time.Sleep(time.Second * 1)
}

func processData(done <-chan bool) {

	for {
		select{
		case <- done:  //true, ou seja, tem dado no canal
			fmt.Println("--- A rotina será encerrada---")
			return

		default:
			fmt.Println("Processando dados")
		}
	}
}
