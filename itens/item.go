package itens

import (
	p "mcronalds/produtos"
)

type Item struct {
	Prod  *p.Produto
	Quant int
}

func (i *Item) CalcularPrecoParcial() float64 {
	return i.Prod.Preco * float64(i.Quant)
}

/*
Retorna um Item com as informações solicitadas.
Se o id não existir para um produto, retorna um Item vazio.
*/
func Criar(id int, quant int) Item {
	produto, _ := p.BuscarId(id)
	if produto == nil {
		return Item{}
	}

	return Item{Prod: produto, Quant: quant}
}
