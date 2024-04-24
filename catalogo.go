package main

import (
	"github.com/JessicaDiehl/loja-digiport-backend/model"
)

func criaCatalogo() []model.Produto { //funcao retorna um array de produto
	produtos := []model.Produto{
		{
			Id:         "1",
			Nome:       "Blusa laura",
			Descricao:  "Cor azul",
			Categoria:  "Blusa",
			Valor:      19.90,
			Quantidade: 2,
			Imagem:     "azul_laura.png",
		},

		{
			Id:         "2",
			Nome:       "Blusa agata",
			Descricao:  "Cor agata",
			Categoria:  "Blusa",
			Valor:      50.00,
			Quantidade: 1,
			Imagem:     "laranja_agata.png",
		},
		{
			Id:         "3",
			Nome:       "Calça Roberta",
			Descricao:  "Calça jeans mom",
			Categoria:  "Calça",
			Valor:      100.00,
			Quantidade: 5,
			Imagem:     "calca_roberta.png",
		},

		{
			Id:         "4",
			Nome:       "Calça Julia",
			Descricao:  "Calça jeans wide",
			Categoria:  "Calça",
			Valor:      100.00,
			Quantidade: 4,
			Imagem:     "calca_julia.png",
		},
		{
			Id:         "5",
			Nome:       "Calça Karol",
			Descricao:  "Calça jeans mom",
			Categoria:  "Calça",
			Valor:      140.00,
			Quantidade: 3,
			Imagem:     "calca_karol.png",
		},
	}

	return produtos
}

func produtosPorNome(nome string) []model.Produto {

	produtosCatalogo := criaCatalogo()

	var produtosFiltrados []model.Produto

	for _, produtos := range produtosCatalogo {
		if produtos.Nome == nome {
			produtosFiltrados = append(produtosFiltrados, produtos)
		}
	}
	return produtosFiltrados
}
