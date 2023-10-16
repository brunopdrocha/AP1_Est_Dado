package main

import (
	"fmt"
	"produtos"
)

func main() {
	var opcao int

	for {
		fmt.Println("Menu:")
		fmt.Println("1 - Cadastrar Produto")
		fmt.Println("2 - Remover Produto")
		fmt.Println("3 - Exibir Todos os Produtos")
		fmt.Println("4 - Buscar Produto por ID")
		fmt.Println("5 - Adicionar Pedido")
		fmt.Println("6 - Expedir Pedido")
		fmt.Println("7 - Exibir Métricas do Sistema")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			fmt.Println("Cadastro de Produto:")
			var nome, descricao string
			var valor float64
			fmt.Print("Nome do Produto: ")
			fmt.Scanln(&nome)
			fmt.Print("Descrição do Produto: ")
			fmt.Scanln(&descricao)
			fmt.Print("Valor do Produto (em reais): ")
			fmt.Scanln(&valor)

			novoProduto = produtos.Produto{
				Nome:      nome,
				Descricao: descricao,
				Valor:     valor,
			}
			novoProduto.CadastrarProduto()
		case 2:
			var id int
			fmt.Print("Digite o ID do Produto que deseja remover: ")
			fmt.Scanln(&id)
			produtos.RemoverProdutoPorID(id)
		case 3:
			produtos.ExibirProdutos()
		case 4:
			var id int
			fmt.Print("Digite o ID do Produto que deseja buscar: ")
			fmt.Scanln(&id)
			produtos.BuscarProdutoPorID(id)
		case 5:
			var id, quantidade int
			fmt.Print("Digite o ID do Produto que deseja adicionar ao pedido: ")
			fmt.Scanln(&id)
			fmt.Print("Digite a quantidade desejada: ")
			fmt.Scanln(&quantidade)
			produtos.AdicionarProdutoAoPedido(id, quantidade)
		case 6:
			produtos.ExpedirPedido()
			fmt.Println("Calculando o valor do pedido...")
			novoPedido.CalcularValorTotal()

			fmt.Println("Informações do Pedido:")
			fmt.Println("Pedido ID:", novoPedido.Id)
			fmt.Println("Delivery:", novoPedido.Delivery)
			fmt.Println("Produtos:")
			for _, produtoPedido := range novoPedido.Produto {
				if produtoPedido.Produto.Id != 0 {
					fmt.Printf("  %s (Quantidade: %d)\n", produtoPedido.Produto.Nome, produtoPedido.Qtde)
				}
			}

			if len(novoPedido.Produto) <= 10 {
				fmt.Println("Pedido expedido com sucesso.")
				novoPedido.Expedir()
			} else {
				fmt.Println("Nenhum pedido aberto para expedição. Limite de produtos excedido.")
			}
		case 7:
			produtos.ExibirMetricasSistema()
			// Exibir métricas do sistema
			numeroTotalProdutos := produtos.NumeroTotalProdutos()
			numeroPedidosEncerrados := produtos.NumeroPedidosEncerrados()
			numeroPedidosEmAndamento := produtos.NumeroPedidosEmAndamento()
			faturamentoTotal := produtos.FaturamentoTotal()

			fmt.Println("Métricas do Sistema:")
			fmt.Println("Número total de produtos cadastrados:", numeroTotalProdutos)
			fmt.Println("Número de pedidos já encerrados:", numeroPedidosEncerrados)
			fmt.Println("Número de pedidos em andamento:", numeroPedidosEmAndamento)
			fmt.Println("Faturamento total até o momento: R$", faturamentoTotal)

		case 0:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
	}
}