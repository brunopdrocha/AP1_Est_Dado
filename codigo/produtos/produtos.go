package produtos

import "fmt"

var TotalProdutos = 0

type Produto struct {
	Id        int
	Nome      string
	Descricao string
	Valor     int
}

func (p *Produto) DefinirId() {
	TotalProdutos++
	var m, d string
	p.Id = TotalProdutos
	fmt.Println("defina o nome do produto")
	fmt.Scanln(&m)
	p.Nome = *&m
	fmt.Println("descricao do Produto")
	fmt.Scanln(&d)
	p.Descricao = *&d

}
func (p*Produto)DefinirValor()  {
	var v int
	fmt.Println("Valor do Produto")
	fmt.Scanln(&v)
	p.Valor = *&v
}
