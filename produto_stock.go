package produto_stock

import (
	//"Stock_Acme/funcoes"
	//"Stock_Acme/repositorioDados"
	"fmt"
)

type Produto struct{
	Identificador            int
	Descricao             string
	CondicaoArmazenamento string
	Categoria             string
}

//var nomeArquivo = "produto"

func (p *Produto ) CadastroProduto(pr []Produto){

	for _, dados := range pr {
		p.Identificador = dados.Identificador
		p.Descricao = dados.Descricao
		p.CondicaoArmazenamento = dados.CondicaoArmazenamento
		p.Categoria = dados.Categoria
	}
	
	/*s :=funcoes.ConverteStructEmString(p)
	r :=funcoes.ConverteDadosEmJson(s)
	repositorioDados.SalavaDadosNoRepositorio(nomeArquivo, r)*/

}

func (p *Produto) VerProdutoCadastrado() string {
	return "IDENTIFICADOR: " + string(p.Identificador) + "DESCRIÇÃO: "+p.Descricao +"\nCONDIÇÃO DE ARMAZENAMENTO: "+p.CondicaoArmazenamento+"\nCATEGORIA: "+p.Categoria
}

func (p Produto) BuscarProdutoPorIdentificador(i int) Produto{
	
	if p.Identificador == i {
	
		return p 
	}
	 return Produto{}

}

func (p Produto) BuscarProdutoPelaCategoria(c string){
	
	if p.Categoria == c {
		fmt.Println("===== Produto encontrado ======== ")
		fmt.Println("Categoria de Produto: ",p.Categoria)
		fmt.Println("Descrição: ",p.Descricao)
		fmt.Println("Condição de armazenamento: ",p.CondicaoArmazenamento)
		return 
	}
	fmt.Println("Produto não encontrado ")

}


