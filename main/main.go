package main

import (
	produto_stock "Stock_Acme"
	 "Stock_Acme/funcoes"
	 "Stock_Acme/repositorioDados"
	"fmt"
)

func main(){
	produto := produto_stock.Produto{}

	produto.CadastroProduto("aaa", "ssss", "eeeee")
	fmt.Println(produto.VerProdutoCadastrado())
	fmt.Println("Descrição: ", produto.Descricao)
	r :=funcoes.ConverteDadosEmJson(produto)

	fmt.Println("Resultado da concersão: ",r)
	repositorioDados.CriaRepositorio()
}