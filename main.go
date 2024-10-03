package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		
		resposta := map[string]string{"mensagem": "Requisição bem-sucedida", "status": "ok"}
		json.NewEncoder(w).Encode(resposta)
	})

	porta := ":9090"
	fmt.Printf("Servidor iniciado na porta %s\n", porta)
	if err := http.ListenAndServe(porta, nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}
