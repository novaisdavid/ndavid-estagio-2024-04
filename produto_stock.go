package produto_stock

import (
	"Stock_Acme/funcoes"
	"Stock_Acme/repositorioDados"
	"fmt"
)

type Produto struct{
	Descricao             string
	CondicaoArmazenamento string
	Categoria             string
}

var nomeArquivo = "produto"

func (p Produto ) CadastroProduto(pr Produto){
	p.Descricao = pr.Descricao
	p.CondicaoArmazenamento = pr.CondicaoArmazenamento
	p.Categoria = pr.Categoria
	fmt.Println("sxxxxxx")
	s :=funcoes.ConverteStructEmString(p)
	r :=funcoes.ConverteDadosEmJson(s)
	repositorioDados.SalavaDadosNoRepositorio(nomeArquivo, r)

}

func (p Produto) VerProdutoCadastrado() string {
	return p.Descricao +" "+p.CondicaoArmazenamento+" "+p.Categoria
}
