package main

import (
	produto_stock "Stock_Acme"
	 "Stock_Acme/funcoes"
	 "Stock_Acme/repositorioDados"
	"fmt"
)

func main(){
	produtos := []produto_stock.Produto{}
	var r []byte

	for i:=0; i<10; i++ {
	
		p := produto_stock.Produto{
			Descricao:            "arroz",
			CondicaoArmazenamento: "pao" ,
			Categoria:            "queijo",
		}
		produtos = append(produtos, p)//.CadastroProduto("arroz"+string(i), "pao"+string(i), "queijo"+string(i))
	}
	//produto.CadastroProduto("arroz", "pao", "queijo")

	for _, prod := range produtos {
		fmt.Println(prod.VerProdutoCadastrado())
		fmt.Println("Descrição: ", prod.Descricao)
		r =funcoes.ConverteDadosEmJson(prod)
		//repositorioDados.SalavaDadosNoRepositorio(r)
	}
	
	

	//fmt.Println("Resultado da concersão: ",r)
	//repositorioDados.CriaRepositorio()
	//repositorioDados.SalavaDadosNoRepositorio(r)
	//repositorioDados.LeDadosDoRepositorio()
}