package main

import (
	"fmt"
	"time"
)

//estrutura q modela a msg do serviço de mensageria
type Message struct {
	Id        int
	Content   string
	CreatedAt time.Time//timestamp
}

func main() {

	//primeiro canal q recebe msgs
	receiver := make(chan Message, 20)

	//segundo canal q é o repositório a ser lido
	sender := make(chan Message, 20)


	//rotina que envia msgs da aplicação para o canal de recebimento
	for i :=1; i <=20; i++ {

		msg := Message{
			Id:i,
			Content: fmt.Sprintf("Mensagem %d recebida.", i),
			CreatedAt: time.Now(),
		}
		//msg sendo enviada ao canal recebedor
		receiver <- msg
		time.Sleep(1 * time.Second)	
	}

	//rotina que transmite esses dados entre canais
	go func ()  {
		for msg := range receiver{
		sender <- msg
		time.Sleep(1 * time.Second)		
	}
	}()

	//rotina q envia as msgs
	go func ()  {
		for msg := range sender {
			fmt.Println(" A mensagem recebida é: ", msg.Content)
			fmt.Println(" Seu timestamp é: ", msg.CreatedAt)
			time.Sleep(1 * time.Second)	
		}
	}()

	time.Sleep(4 * time.Second)	
	close(receiver)
	close(sender)

}