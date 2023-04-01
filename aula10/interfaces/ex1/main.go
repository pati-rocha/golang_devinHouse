//interfaces - são tipo "conjuntos" de métodos e/ou assinaturas (ações)

package main

import "fmt"

//declaração básica/padrão de interfaces
type Contador interface {
	Adicionar(i int) int //métodos
	PegarValor() int
}

//implementação básica da interface
type Status struct {
	Valor int //atributos
}

func (s *Status) Adicionar(i int) int {
	z := s.Valor + i
	return z
}

func ( s *Status) PegarValor() int {
	return s.Valor
}

func IniciarMonitoramento(c Contador) {

	c.Adicionar(2)
	fmt.Println(c.PegarValor())

}

func main() {
	s := Status{1}
	fmt.Println(s)
	fmt.Println(s.Adicionar(1))

	IniciarMonitoramento(&s)
}