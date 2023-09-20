package produtos

import "fmt"

var TotalProdutos = 0

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Valor     int
}

func NewProduto(nome string) *Produto {
	TotalProdutos++
	i.Id = TotalProdutos
	fmt.Println("defina o Valor do produto")
	fmt.Scanln(TotalProdutos)
	p.Valor= Valor
	fmt.Println("defina o nome do produto")
	fmt.Scanln(&p)
	p.Nome = Nome
	return &Produto{Id: TotalProdutos, Nome: nome}
}

func (p *Produto) DefinirDesc() {
	fmt.Println("defina o nome do produto")
	fmt.Scanln(&p)
	p.Nome = p.Nome
}
func (p *Produto) DefinirValor() {
	fmt.Println("defina o Valor do produto")
	fmt.Scanln(&p)
	p.Valor= p.Valor
}