// main.go

package main

import (
	"Ap1/codigo/produtos"
	"fmt"
)

func main() {
	var opcao int

	fmt.Println("Bem-vindo à CLI do Sistema de Pedidos!")

	for {
		fmt.Println("Menu:")
		fmt.Println("1 - Cadastrar Produto")
		fmt.Println("2 - Remover Produto")
		fmt.Println("3 - Exibir Todos os Produtos")
		fmt.Println("4 - Buscar Produto por ID")
		fmt.Println("5 - Adicionar Produto ao Pedido")
		fmt.Println("6 - Expedir Pedido")
		fmt.Println("7 - Exibir Métricas do Sistema")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			produto := produtos.Produto{}
			produto.CadastrarProduto()
		case 2:
			produtos.RemoverProduto()
			fmt.Scanln(&opcao)
		case 3:
			produtos.ExibirProdutos()
		case 4:
			produto := produtos.Produto{}
			produto.BuscaProduto()
		case 5:
			pedido := produtos.Pedido{}
			pedido.AdicionaPedido()
			fmt.Scanln(&opcao)

		case 6:
			produtos.ExpedirPedido()

		case 7:
			produtos.ExibirMetricasSistema()

		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}
