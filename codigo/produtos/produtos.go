package produtos

import "fmt"

var TotalProdutos = 0
var ListaProduto [10]*Produto
var TotalPedidos = 0
var ListaPedido [1000]*Pedido
var numeroPedidosEmAndamento int
var Tudo float64

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Valor     float64
}

type ProdutoPedido struct {
	Produto Produto
}

type Pedido struct {
	Id         int
	Delivery   bool
	Produto    [10]ProdutoPedido
	ValorTotal float64
}

func (p *Produto) CadastrarProduto() {
	TotalProdutos++
	var m, d string
	var v float64
	p.Id = TotalProdutos
	fmt.Println("Nome do produto:")
	fmt.Scanln(&m)
	p.Nome = m
	fmt.Println("Descrição do produto:")
	fmt.Scanln(&d)
	p.Descricao = d
	fmt.Println("Valor do produto:")
	fmt.Scanln(&v)
	p.Valor = v

	ListaProduto[TotalProdutos-1] = p
}

func RemoverProduto() {
	fmt.Println("Digite o ID do Produto que deseja remover:")
	var id int
	fmt.Scan(&id)
	for i, produto := range ListaProduto {
		if produto != nil && produto.Id == id {
			ListaProduto[i] = nil
			fmt.Println("Produto removido com sucesso")
			break
		}
	}
}

func ExibirProdutos() {
	if TotalProdutos == 0 {
		fmt.Println("Nenhum produto cadastrado.")
		return
	}

	for _, produto := range ListaProduto {
		if produto != nil && produto.Id != 0 {
			fmt.Println("ID:", produto.Id)
			fmt.Println("Nome:", produto.Nome)
			fmt.Println("Descrição:", produto.Descricao)
			fmt.Println("Valor:", produto.Valor)
			fmt.Println("--------------")
		}
	}
}

func (p *Produto) BuscaProduto() {
	fmt.Print("Informe o ID do produto que deseja exibir: ")
	var id int
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Erro ao ler ID do produto")
		return
	}

	found := false
	for _, produto := range ListaProduto {
		if produto != nil && produto.Id == id {
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

func (p *Pedido) AdicionaPedido() {
	if TotalPedidos >= len(ListaPedido) {
		fmt.Println("Atingiu o limite máximo de pedidos. Não é possível adicionar mais pedidos.")
		return
	}

	var maisProdutos bool = true
	indiceProduto := 0

	for maisProdutos {
		var produtoID int
		fmt.Print("Digite o ID do produto que deseja adicionar: ")
		fmt.Scan(&produtoID)

		if produtoID >= 1 && produtoID <= TotalProdutos {
			produtoPedido := ProdutoPedido{Produto: *ListaProduto[produtoID-1]}
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
		fmt.Println("1-Delivery/0-Retirada")
		fmt.Scanln(&p.Delivery)

	}

	p.CalcularValorTotal()
}

func (pp *ProdutoPedido) CalcularPreco() float64 {
	return float64(pp.Produto.Valor)
}

func (p *Pedido) CalcularValorTotal() {
	var total float64
	for _, produtoPedido := range p.Produto {
		total += produtoPedido.CalcularPreco()
	}
	var frete float64 = 10.0
	p.ValorTotal = total + frete
	fmt.Println("Valor total do pedido R$:", total)
	fmt.Println("Valor total do pedido com frete (R$ 10,00) R$:", p.ValorTotal)
	fmt.Scanln(&p.Delivery)

	if p.Delivery == true {
		Tudo += p.ValorTotal
	} else {
		Tudo += total
	}

}

func ExpedirPedido() {
	for i, pedido := range ListaPedido {
		if pedido != nil && pedido.Id != 0 && pedido.ValorTotal > 0 {
			fmt.Println("Expedindo pedido ID:", pedido.Id)
			fmt.Println("Delivery:", pedido.Delivery)
			fmt.Println("Produtos:")
			for _, produtoPedido := range pedido.Produto {
				if produtoPedido.Produto.Id != 0 {
					fmt.Println(produtoPedido.Produto.Nome)
				}
			}
			fmt.Println("Valor Total: R$", pedido.ValorTotal)

			// Atualize a variável produtopa para refletir a expedição do pedido

			ListaPedido[i] = nil
			break
		}
	}
	fmt.Println("Pedido expedido.")
}

func pedidoExpedido() bool {
	numeroPedidosEncerrados := 0
	for _, pedido := range ListaPedido {
		if pedido != nil && pedido.Id == 0 && pedido.ValorTotal > 0 {
			numeroPedidosEncerrados++
		}
	}
	return numeroPedidosEncerrados > 0
}

func ExibirMetricasSistema() {
	numeroTotalProdutos := 0
	for _, produto := range ListaProduto {
		if produto != nil && produto.Id != 0 {
			numeroTotalProdutos++
		}
	}

	faturamentoTotal := 0.0
	for _, pedido := range ListaPedido {
		if pedido != nil && pedido.Id == 0 {
			faturamentoTotal += pedido.ValorTotal
		}
	}

	fmt.Println("Número total de produtos cadastrados:", numeroTotalProdutos)
	fmt.Println("Faturamento total até o momento: R$", faturamentoTotal)
}
