package produtos

import (
    m "mcronalds/metricas"
    "strings"

)



type NoProduto struct {
    Produto Produto
    Next    *NoProduto
}

type ListaProdutos struct {
    Head *NoProduto
    Tail *NoProduto
}

var TotalProdutosJaCadastrados = 0
var ListaDeProdutos ListaProdutos

func tentarCriar(nome, descricao string, preco float64, id int) *Produto {
    if id != -1 {
        if _, idProcurado := BuscarId(id); idProcurado != -1 {
            return nil
        }
    }
    return &Produto{Id: id, Nome: nome, Descricao: descricao, Preco: preco}
}

func AdicionarUnico(nome, descricao string, preco float64, id int) int {
    novoProduto := tentarCriar(nome, descricao, preco, id)
    if novoProduto == nil {
        return -3 // Erro ao criar o produto
    }

    ListaDeProdutos.AdicionarProduto(*novoProduto)
    m.M.SomaProdutosCadastrados(1)
    TotalProdutosJaCadastrados++
    return TotalProdutosJaCadastrados
}

func BuscarId(id int) (*Produto, int) {
    atual := ListaDeProdutos.Head
    indice := 0

    for atual != nil {
        if atual.Produto.Id == id {
            return &atual.Produto, indice
        }
        atual = atual.Next
        indice++
    }

    return nil, -1
}

func BuscarNome(comecaCom string) ([]Produto, int) {
    var produtosEncontrados []Produto
    atual := ListaDeProdutos.Head

    for atual != nil {
        if strings.HasPrefix(atual.Produto.Nome, comecaCom) {
            produtosEncontrados = append(produtosEncontrados, atual.Produto)
        }
        atual = atual.Next
    }

    return produtosEncontrados, len(produtosEncontrados)
}

func Exibir() {
    atual := ListaDeProdutos.Head

    for atual != nil {
        atual.Produto.Exibir()
        atual = atual.Next
    }
}

func ExibirPorNome() {
    produtosOrdenados := make([]Produto, TotalProdutosJaCadastrados)
    atual := ListaDeProdutos.Head

    for i := 0; atual != nil; i++ {
        produtosOrdenados[i] = atual.Produto
        atual = atual.Next
    }

    for i := 0; i < TotalProdutosJaCadastrados-1; i++ {
        minIndex := i
        for j := i + 1; j < TotalProdutosJaCadastrados; j++ {
            if strings.ToLower(produtosOrdenados[j].Nome) < strings.ToLower(produtosOrdenados[minIndex].Nome) {
                minIndex = j
            }
        }
        if minIndex != i {
            produtosOrdenados[i], produtosOrdenados[minIndex] = produtosOrdenados[minIndex], produtosOrdenados[i]
        }
    }

    for i := 0; i < TotalProdutosJaCadastrados; i++ {
        produtosOrdenados[i].Exibir()
    }
}

func Excluir(id int) int {
    atual := ListaDeProdutos.Head
    anterior := &NoProduto{}

    if atual == nil {
        return -2 // Lista vazia
    }

    if atual.Produto.Id == id {
        ListaDeProdutos.Head = atual.Next
        m.M.SomaProdutosCadastrados(-1)
        TotalProdutosJaCadastrados--
        return 0 // Exclusão bem-sucedida
    }

    for atual != nil {
        if atual.Produto.Id == id {
            anterior.Next = atual.Next
            m.M.SomaProdutosCadastrados(-1)
            TotalProdutosJaCadastrados--
            return 0 // Exclusão bem-sucedida
        }
        anterior = atual
        atual = atual.Next
    }

    return -1 // Produto não encontrado
}


