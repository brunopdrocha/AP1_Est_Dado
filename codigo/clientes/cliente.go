package clientes

import (
	p "Ap1/codigo/produtos"
	"fmt"
)

func Cli() {
	var p1, p2, p3, p4, p5,p6,p7,p8,p9,p10 p.Produto
	p1 = p.Produto{Nome: "prod1", Descricao: "descricao"}
	p1.DefinirId()
	p1.DefinirValor()

	p2 = p.Produto{Nome: "prod2", Descricao: "descricao"}
	p2.DefinirId()
	p2.DefinirValor()
	p3 = p.Produto{Nome: "prod3", Descricao: "descricao"}
	p3.DefinirId()
	p3.DefinirValor()
	p4 = p.Produto{Nome: "prod4", Descricao: "descricao"}
	p4.DefinirId()
	p4.DefinirValor()
	p5 = p.Produto{Nome: "prod5", Descricao: "descricao"}
	p5.DefinirId()
	p5.DefinirValor()
	p6 = p.Produto{Nome: "prod6", Descricao: "descricao"}
	p6.DefinirId()
	p6.DefinirValor()
	p7 = p.Produto{Nome: "prod7", Descricao: "descricao"}
	p7.DefinirId()
	p7.DefinirValor()
	p8 = p.Produto{Nome: "prod8", Descricao: "descricao"}
	p8.DefinirId()
	p8.DefinirValor()
	p9 = p.Produto{Nome: "prod9", Descricao: "descricao"}
	p9.DefinirId()
	p9.DefinirValor()
	p10 = p.Produto{Nome: "prod10", Descricao: "descricao"}
	p10.DefinirId()
	p10.DefinirValor()

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(p4)
	fmt.Println(p5)
	fmt.Println(p6)
	fmt.Println(p7)
	fmt.Println(p8)
	fmt.Println(p9)
	fmt.Println(p10)

	fmt.Println(p.TotalProdutos)
}
