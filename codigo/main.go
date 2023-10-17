package main

import (
    "Ap1/codigo/produtos"
    "fmt"
)

func main() {
    var pedidoAberto int
    var pedidofechado int
    var qtdpedidos int
    var opcao int

    fmt.Println("Bem-vindo à CLI do McRonald’s")

    for {
        fmt.Println("Menu:")
        fmt.Println("1 - Cadastrar Produto McRonald’s")
        fmt.Println("2 - Remover Produto McRonald’s")
        fmt.Println("3 - Exibir Todos os Produtos McRonald’s")
        fmt.Println("4 - Buscar Produto por ID")
        fmt.Println("5 - Adicionar Produto ao Pedido McRonald’s")
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
			fmt.Scanln(&opcao)
        case 4:
            produto := produtos.Produto{}
            produto.BuscaProduto()
        case 5:
            pedido := produtos.Pedido{}
			pedido.AdicionaPedido()
            if opcao == 5 {
                pedidoAberto++
                qtdpedidos++
            }

        case 6:
            produtos.ExpedirPedido()
            if opcao == 6 {
				if pedidoAberto <= 0 {
					break
				}
                pedidoAberto--
                pedidofechado++
            }
        case 7:
            fmt.Println("O numero de pedidos PENDENTES: ", pedidoAberto)
            fmt.Println("O numero de pedidos CONCLUIDOS: ", pedidofechado)
            fmt.Println("O numero de pedidos RECEBIDOS: ", qtdpedidos)
            fmt.Println("O Faturamento foi de  ", produtos.Tudo)
        case 0:
            fmt.Println("Saindo do McRonald’s...")
            return
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
