package main

/*ter uma database em memória sobre carros e ser capaz de filtrá-la por marcas*/

import "fmt"


//struct que define entidade Carro
type Carro struct{
	Marca string
	Modelo string
	Ano int
}

func filtrarCarrosPorMarca( carros []Carro, marca string) []Carro {
	CarrosFiltrados := []Carro{}

	for _, c := range carros {
		if c.Marca == marca {
			CarrosFiltrados = append(CarrosFiltrados, c)
		}
	}
	return CarrosFiltrados

}
func main(){
	carros := []Carro {
		{Marca:"Toyota", Modelo: "Corolla", Ano: 2015},	
		{Marca:"Fiat", Modelo: "Mobi", Ano: 2020},	
		{Marca:"Volkswagem", Modelo: "Fusca", Ano: 1975},	
		{Marca:"Honda", Modelo: "Civic", Ano: 2016},	
		{Marca:"Fiat", Modelo: "Uno", Ano: 2009},	
		{Marca:"Chevrolet", Modelo: "Opala", Ano: 1989},	
		{Marca:"Ferrari", Modelo: "Enzo", Ano: 2002},	
	}

	f := filtrarCarrosPorMarca(carros, "Chevrolet")

	for _, carro := range f{
		fmt.Println(carro)
	}

}