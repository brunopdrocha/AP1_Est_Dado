package produtos

import "fmt"

var TotalProdutos = 0
var ListaProduto [50]*Produto
var TotalPedidos = 0
var ListaPedido [1000]*Pedido
var FaturamentoTotal = 0

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Valor     float64
}

type ProdutoPedido struct {
	Produto Produto
	Qtde    int
}

type Pedido struct {
	Id         int
	Delivery   bool
	Produto    [10]ProdutoPedido
	ValorTotal float64
}

//Produto

// Definicao dos Produtos
func (p *Produto) CadastrarProduto() {

	TotalProdutos++
	var m, d string
	var v float64
	p.Id = TotalProdutos
	fmt.Println("defina o nome do produto")
	fmt.Scanln(&m)
	p.Nome = m
	fmt.Println("descricao do Produto")
	fmt.Scanln(&d)
	p.Descricao = d
	fmt.Println("Valor do Produto")
	fmt.Scanln(&v)
	p.Valor = v

	// Adiciona o produto à ListaProduto (slice)
	ListaProduto[TotalProdutos-1] = p
}

//Remover Produtos
func (p *Produto) RemoverProduto() {
	// Remover Produto da lista ordenada

	fmt.Println("Digite o ID do Produto")
	var id int
	fmt.Scan(&id)

	//Caso Base
	if p == nil || p.Id == 0 {
		fmt.Println("Nao possui nenhum produto")
	}

	if p.Id == id {
		*p = Produto{}
		fmt.Println("Produto foi removido com sucesso")
	} else {
		fmt.Println("Produto não encontrado")
	}

}

//Exibir Todos os Produtos
func ExibirProdutos() {
	fmt.Println("Lista de Produtos Cadastrados:")
	for _, produto := range ListaProduto {
		if produto.Id != 0 {
			fmt.Println("ID:", produto.Id)
			fmt.Println("Nome:", produto.Nome)
			fmt.Println("Descrição:", produto.Descricao)
			fmt.Println("Valor:", produto.Valor)
			fmt.Println("--------------")
		}
	}
}

//Exibir Determinado Produto pelo ID do Produto
func (p *Produto) BuscaProduto() {
	fmt.Println("Informe o ID do produto que deseja exibir: ")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Erro ao ler ID do produto")
		return
	}

	// Buscar o produto pelo ID
	var found bool
	for _, produto := range ListaProduto {
		if produto.Id == id {
			fmt.Println("Nome:", produto.Nome)
			fmt.Println("Descrição:", produto.Descricao)
			fmt.Println("Valor:", produto.Valor)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Produto não encontrado")
	}

}

//Calcula valor do Pedido
func (pp *ProdutoPedido) CalcularPreco() float64 {
	floatQtde := float64(*&pp.Qtde)
	return float64(pp.Produto.Valor * floatQtde)
}

//Sera adicionado determinado pedido
func (p *Pedido) AdicionaPedido() {
	if TotalPedidos >= len(ListaPedido) {
		fmt.Println("Atingiu o máximo de pedidos. Não é possível adicionar mais pedidos.")
		return
	}

	var maisProdutos bool = true

	indiceProduto := 0

	for maisProdutos {
		// Código para adicionar produtos ao pedido
		var produtoID, quantidade int
		fmt.Print("Digite o ID do produto que deseja adicionar: ")
		fmt.Scan(&produtoID)
		fmt.Print("Digite a quantidade desejada: ")
		fmt.Scan(&quantidade)

		// Adicionar o ProdutoPedido ao próximo índice disponível em Produtos
		if produtoID >= 1 && produtoID <= TotalProdutos {
			produtoPedido := ProdutoPedido{Produto: *ListaProduto[produtoID-1], Qtde: quantidade}
			p.Produto[indiceProduto] = produtoPedido
			indiceProduto++

			if indiceProduto >= len(p.Produto) {
				fmt.Println("Você atingiu o limite de produtos para este pedido (10 itens).")
				break
			}
		} else {
			fmt.Println("Produto não encontrado. Por favor, digite um ID de produto válido.")
		}

		fmt.Print("Deseja adicionar mais produtos? (1-Sim / 2-Não): ")
		var opcao int
		fmt.Scan(&opcao)
		maisProdutos = opcao == 1
	}

	p.CalcularValorTotal()
}

// CalcularValorTotal atualizado para incluir a taxa de entrega
func (p *Pedido) CalcularValorTotal() {
	var total float64
	for _, produtoPedido := range p.Produto {
		total += produtoPedido.CalcularPreco()
	}
	var frete float64 = 10.0 // Taxa de entrega
	p.ValorTotal = total + frete
	fmt.Println("Valor total do pedido R$:", p.ValorTotal)
}

// ExpedirPedido verifica o primeiro pedido ainda aberto, exibe suas informações e adiciona ao faturamento
func ExpedirPedido() {
	var pedidoExpedido bool

	// Encontrar o primeiro pedido ainda aberto
	for _, pedido := range ListaPedido {
		if pedido != nil && !pedidoExpedido {
			if pedido.Id != 0 && pedido.ValorTotal > 0 {
				// Exibir informações do pedido
				fmt.Println("Pedido ID:", pedido.Id)
				fmt.Println("Delivery:", pedido.Delivery)
				fmt.Println("Produtos:")
				for _, produtoPedido := range pedido.Produto {
					if produtoPedido.Produto.Id != 0 {
						fmt.Printf("  %s (Quantidade: %d)\n", produtoPedido.Produto.Nome, produtoPedido.Qtde)
					}
				}
				fmt.Println("Valor Total: R$", pedido.ValorTotal)

				// Marcar o pedido como expedido
				pedido.Id = 0
				pedido.Delivery = false
				pedido.Produto = [10]ProdutoPedido{}
				pedido.ValorTotal = 0

				fmt.Println("Pedido expedido com sucesso.")
				pedidoExpedido = true
			}
		}
	}
	// Se nenhum pedido foi expedido, exibir uma mensagem
	if !pedidoExpedido {
		fmt.Println("Nenhum pedido aberto para expedição.")
	}
}

// ExibirMetricasSistema exibe métricas do sistema
func ExibirMetricasSistema() {
	// Número total de produtos cadastrados
	numeroTotalProdutos := TotalProdutos

	// Número de pedidos já encerrados
	numeroPedidosEncerrados := 0
	for _, pedido := range ListaPedido {
		if pedido != nil && pedido.Id == 0 {
			numeroPedidosEncerrados++
		}
	}

	// Número de pedidos em andamento
	numeroPedidosEmAndamento := 0
	for _, pedido := range ListaPedido {
		if pedido != nil && pedido.Id != 0 {
			numeroPedidosEmAndamento++
		}
	}

	// Faturamento total até o momento
	faturamentoTotal := FaturamentoTotal

	// Exibir métricas do sistema
	fmt.Println("Número total de produtos cadastrados:", numeroTotalProdutos)
	fmt.Println("Número de pedidos já encerrados:", numeroPedidosEncerrados)
	fmt.Println("Número de pedidos em andamento:", numeroPedidosEmAndamento)
	fmt.Println("Faturamento total até o momento: R$", faturamentoTotal)
}
