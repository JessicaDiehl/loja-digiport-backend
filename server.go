package main

import (
	"encoding/json"
	"net/http"

	"github.com/JessicaDiehl/loja-digiport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {

	queryParam := r.URL.Query()

	nome := queryParam.Get("nome")

	produtos := criaCatalogo()

	if nome != "" {
		filteredProdutos := []model.Produto{}
		for _, p := range produtos {
			if p.Nome == nome {
				filteredProdutos = append(filteredProdutos, p)
			}
		}
		produtos = filteredProdutos
	}

	json.NewEncoder(w).Encode(produtos)
}
