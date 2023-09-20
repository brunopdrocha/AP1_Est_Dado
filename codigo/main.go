package main

import (
	c "Ap1/codigo/clientes"
	p "Ap1/codigo/produtos"
	"fmt"
)

func main() {
	for {
		var op int

		op = 0
		fmt.Println("escolha a função que deseja")
		fmt.Println("1- voltando para a interface")
		fmt.Println("2- Add produtos Até 10 produtos ")
		fmt.Scanln(&op)
		switch op {
		case 1:
			{
				{
					fmt.Println("VOLTANDO PRA INTERFACE")
				}
			}

		case 2:
			{
				c.Cli()
				fmt.Println("ADD Produto")
				fmt.Println(p.TotalProdutos)

			}
		}

	}
}
