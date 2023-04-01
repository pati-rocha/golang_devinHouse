package main

import "fmt"

//pegar os preços dos itens

//interface = contrato
type Item interface {
	GetPrice() float64 //assinatura -> é o nome que identifica o método
}

//structs - define quais atributos terá a entidade ou a estrutura de dados
type Hamburguer struct {
	Name string
	Price float64
}

type Batata struct {
	Size string
	Price float64
}

type Drink struct {
	Name string
	Price float64
}

//método - é a conexão via implementação da interface e da struct
func (h Hamburguer) GetPrice() float64 {
	return h.Price
}
func (b Batata) GetPrice() float64 {
	return b.Price
}
func (d Drink) GetPrice() float64 {
	return d.Price
}	
	
//função geral

func CalculateTotal(itens []Item) float64 {

	var total float64
	
	for _, item := range itens {

		total += item.GetPrice()
	}
	return total
}

func main() {
	itens := []Item{
		Hamburguer{Name: "X-Salada", Price:25},
		Hamburguer{Name: "X-Bacon", Price:30},
		Batata{Size: "Grande", Price:18},
		Drink{Name:"Dolly", Price: 3},
	}
	fmt.Println("Preço Total do Pedido: ", CalculateTotal(itens))
}