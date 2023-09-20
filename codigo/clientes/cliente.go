package clientes

import (
	"fmt"
	p "Ap1/codigo/produtos"
)

func Cli() {
	var p1, p2 p.Produto
	p1 = p.Produto{Nome: "prod1"}
	p1.DefinirId()
	p2 = p.Produto{Nome: "prod2"}
	p2.DefinirId()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p.TotalProdutos)
}