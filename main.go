// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	produtosCatalogo := criaCatalogo()
	//fmt.Printf("Esse é o catalogo da loja: %+v\n\n", produtosCatalogo)

	fmt.Println("Esse é o catálogo da loja: ")

	for _, produtos := range produtosCatalogo {
		fmt.Printf("ID: %s\n", produtos.Id)
		fmt.Printf("Nome: %s\n", produtos.Nome)
		fmt.Printf("Descrição: %s\n", produtos.Descricao)
		fmt.Printf("Categoria: %s\n", produtos.Categoria)
		fmt.Printf("Valor: %.2f\n", produtos.Valor)
		fmt.Printf("Quantidade: %d\n", produtos.Quantidade)
		fmt.Printf("Imagem: %s\n", produtos.Imagem)
		fmt.Println("-------------------------------------------")
	}
}
