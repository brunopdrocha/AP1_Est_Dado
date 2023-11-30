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
    novoID := TotalProdutosJaCadastrados+ 1

    novoProduto := tentarCriar(nome, descricao, preco, novoID)
    if novoProduto == nil {
        return -3 // Erro ao criar o produto
    }

    ListaDeProdutos.AdicionarProduto(*novoProduto)
    m.M.SomaProdutosCadastrados(1)
    TotalProdutosJaCadastrados++
    return novoProduto.Id
}

func obterUltimoID() int {
    if ListaDeProdutos.Tail != nil {
        return ListaDeProdutos.Tail.Produto.Id
    }
    return 0
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

func bubbleSortPorNome(produtos *[]Produto, total int) {
    trocou := true

    for trocou {
        trocou = false
        total--
        for i := 0; i < total; i++ {
            if strings.ToLower((*produtos)[i].Nome) > strings.ToLower((*produtos)[i+1].Nome) {
                (*produtos)[i], (*produtos)[i+1] = (*produtos)[i+1], (*produtos)[i]
                trocou = true
            }
        }
    }
}

func ExibirPorNome() {
    produtosOrdenados := make([]Produto, TotalProdutosJaCadastrados)
    atual := ListaDeProdutos.Head

    for i := 0; atual != nil; i++ {
        produtosOrdenados[i] = atual.Produto
        atual = atual.Next
    }

    bubbleSortPorNome(&produtosOrdenados, TotalProdutosJaCadastrados)

    for i := 0; i < TotalProdutosJaCadastrados; i++ {
        produtosOrdenados[i].Exibir()
    }
}

func Excluir(id int) (int, int) {
    atual := ListaDeProdutos.Head
    anterior := &NoProduto{}

    if atual == nil {
        return -2, -1 // Lista vazia
    }

    if atual.Produto.Id == id {
        ListaDeProdutos.Head = atual.Next
        reajustarIDs(id) // Reajustar IDs após a remoção
        m.M.SomaProdutosCadastrados(-1)
     

        // Se a lista ficar vazia após a remoção, ajuste o Tail para nil
        if ListaDeProdutos.Head == nil {
            ListaDeProdutos.Tail = nil
        }

        return 0, id // Exclusão bem-sucedida
    }

    for atual != nil {
        if atual.Produto.Id == id {
            anterior.Next = atual.Next
            reajustarIDs(id) // Reajustar IDs após a remoção
            m.M.SomaProdutosCadastrados(-1)
          
            // Se o elemento removido for o último, ajuste o Tail para o anterior
            if atual.Next == nil {
                ListaDeProdutos.Tail = anterior
            }

            return 0, id // Exclusão bem-sucedida
        }
        anterior = atual
        atual = atual.Next
    }

    return -1, -1 // Produto não encontrado
}

func reajustarIDs(idRemovido int) {
    atual := ListaDeProdutos.Head
    for atual != nil {
        if atual.Produto.Id > idRemovido {
            atual.Produto.Id--
        }
        atual = atual.Next
    }
}
