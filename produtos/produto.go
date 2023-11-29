package produtos

import "fmt"

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Preco     float64
}


func (p Produto) Exibir() {
	fmt.Println("\nProduto", p.Id)
	fmt.Println(p.Nome)
	fmt.Println(p.Descricao)
	fmt.Printf("Preço: R$ %.2f\n", p.Preco)
}

func CriarProduto(nome, descricao string, preco float64, id int) Produto {
	p := Produto{Nome: nome, Descricao: descricao, Preco: preco}
	if id != -1 {
		p.Id = id
	}
	return p
}

func (lp *ListaProdutos) AdicionarProduto(p Produto) {
	novoNo := &NoProduto{Produto: p, Next: nil}
	if lp.Head == nil {
		lp.Head = novoNo
		lp.Tail = novoNo
	} else {
		lp.Tail.Next = novoNo
		lp.Tail = novoNo
	}
}

func AtualizarPreco(lista *ListaProdutos, id int, novoPreco float64) int {
	produtoEncontrado, indice := BuscarProduto(lista, id)
	if indice == -1 {
		return -1 // Produto não encontrado
	}

	// Atualiza apenas o preço do produto encontrado
	produtoEncontrado.Preco = novoPreco

	return 0 // Atualização bem-sucedida
}

func BuscarProduto(lista *ListaProdutos, id int) (*Produto, int) {
	atual := lista.Head
	indice := 0

	for atual != nil {
		if atual.Produto.Id == id {
			return &atual.Produto, indice
		}
		atual = atual.Next
		indice++
	}

	return nil, -1 // Produto não encontrado
}
