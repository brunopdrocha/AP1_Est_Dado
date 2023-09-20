package main

import (
	"fmt"
	c "Ap1/codigo/clientes"
	p "Ap1/codigo/produtos"
)

func main() {

	c.Cli()
	fmt.Println("--------------")
	fmt.Println(p.TotalProdutos)
	fmt.Println(p.DefinirNome)

}