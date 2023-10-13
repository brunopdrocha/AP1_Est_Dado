package produtos

import "fmt"

var TotalProdutos = 0
var ListaProduto [50]*Produto

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Valor     int
}

func (p *Produto) CadastrarProduto() {
	// Definicao dos Produtos
	TotalProdutos++
	var m, d string
	var v int
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

func (p *Produto) ExibirProduto() {
	// Devemos exibir as informações do produto
	fmt.Print("Informe o ID do produto que deseja exibir: ")
	var id int
	_, err := fmt.Scan(&id)
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
