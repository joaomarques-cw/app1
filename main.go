package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olá! Esta é uma aplicação simples para implantação no Kubernetes/ArgoCD.")
	})

	porta := ":8080"
	fmt.Printf("Servidor iniciado na porta %s\n", porta)
	if err := http.ListenAndServe(porta, nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
