package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func calculaeRetornaMensagem(texto string) string {

	x := 0.0001
	i := 0

	for ; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return texto
}

func calculaeRetornaMensagemHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 nao encontrado.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Metodo nao suportado.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, calculaeRetornaMensagem("Code Education Rocks!"))
}

func main() {

	http.HandleFunc("/", calculaeRetornaMensagemHandler)

	log.Println("Servidor Web iniciado na porta 8000.")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
