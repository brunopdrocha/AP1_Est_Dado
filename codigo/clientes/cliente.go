package clientes

import (
	"Ap1/codigo/produtos"
	"fmt"
)

func Cli() {
	var escolha int
	fmt.Println("Quantos itens você deseja inserir? 0-10")
	fmt.Scanln(&escolha)

	if escolha < 0 || escolha > 10 {
		fmt.Println("Escolha um número entre 0 e 10.")
		return
	}

	var listaProdutos []produtos.Produto

	for i := 0; i < escolha; i++ {
		produto := produtos.Produto{Nome: "prod1", Descricao: "descricao"}
		produto.DefinirId()
		produto.DefinirValor()
		listaProdutos = append(listaProdutos, produto)
		fmt.Println(produto)
	}

	fmt.Println("Total de produtos inseridos:", len(listaProdutos))
	fmt.Println(listaProdutos)
}
