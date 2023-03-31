package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {

		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "Post resquest foi bem sucedido! \n")
	name := r.FormValue("name")
	city := r.FormValue("city")
	uf := r.FormValue("uf")
	

	fmt.Fprintf(w, "Nome = %s\n", name)
	fmt.Fprintf(w, "Cidade = %s\n", city)
	fmt.Fprintf(w, "Estado = %s\n", uf)

}

// função principal do script
func main() {

	//servidor de arquivos estáticos do webserver
	fileServer := http.FileServer(http.Dir("./static"))

	//root handle
	http.Handle("/", fileServer)

	//form handle
	http.HandleFunc("/form", formHandler)

	fmt.Print("Iniciando web-servidor na porta 8080. \n")

	//tratamento de erro
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
