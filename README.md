# AP1_Est_Dado
GRUPO:
Mateus Norcia 
Bruno Xavier
Bruno Pilão
Ian Amoedo
Anotações para o Projeto de AP1 para uma hamburgueria

# Definicao de Produto
  Devemos possuir um contador de produtos inicializando no valor 0

    EXEMPLO 
    var TotalProdutos = 0

    Produto
    id int
    Nome string
    Descricao string
    Preco float64

    Definir id()
      TotalProdutos ++
      //Alterar o id do Produto


# Cadastro de Produto

  Devemos armazenar os produtos em uma lista com um limite de 50 elementos

    EXEMPLO DO PRODUTO
    
    ListaProdutos
    -Maximo de 50 itens
    -Lista Sequencial

    Ações
    -Adicionar O(n):
      -Adicione o produto no final da lista
      -Assumir que os produtos são unicos
                      ou
      -Adicione o produto no final da lista
      -Não adiciono se nome ja existir
    -Remover O (n)

    -BuscarId O(log n)
    -Buscar Nome
    -Exibir

  # Cadastro de Pedido

  Devemos armazenar a quantidade de pedidos possuindo 10 produtos

    EXEMPLO DO PEDIDO

    Pedido

    Id int 
    Delivery bool deve (false comer no locar / true realizar entrega) 
    Produtos [10] Produto Pedido
    DataHora ?
    Valor total float64

    CalculoValorTotal()
    //Soma os valores realizado pelo pedido do cliente

    DefinirId()
    //Definicao de cada Produto

    ProdutoPedido()
      P Produto(Baseado a classe STRUCT Produto)
      qtde int

      Calcular Preco()
      //Calcular preco de cada Produto


  # Fila dos Pedidos

  Devemos controlar a fila dos pedidos suportando ate 1000 pedidos

    EXEMPLO DA FILA

    Obs: Devemos puxar o valores do pedido para a fila
    
    -Lista no maximo de 1000 pedidos
    -Fila
    -Adicionar 
    -Expandir
      -Faturamento (Faturamento deve ser atualizado)
      -Tempo medio

      
      

  
  # Requisitos EXTRAS 

   Há três requisitos que os sócios do McRonald’s gostariam muito que fossem implementados já no MVP, mas que não estão no escopo do serviço:

-  Cadastro de produtos em lote: os sócios gostariam de, com um único comando na CLI, poder passar os dados de vários produtos ao mesmo tempo, para que o sistema armazenasse todas as informações sem precisar inserir a opção de cadastro múltiplas vezes;
-  Leitura dos produtos em arquivo de texto: caso os produtos possam ser lidos a partir de um arquivo .csv, isso agilizaria muito o processo de testes e validação do sistema;
-  Exibir os pedidos ainda em aberto: exibe na tela todos os pedidos que estão em andamento, na ordem em que foram abertos.
    

