package lote

import (
	produto_stock "Stock_Acme"
	"Stock_Acme/funcoes"
	"Stock_Acme/repositorioDados"
)

type Lote struct{

	IdentificadorLote string
	Prateleira        string
	Corredor          string
	Quantidade          int
	DataDeEntrada     string
	DataDeValidade    string
	Produto           produto_stock.Produto
}

var nomeArquivo = "lote"

func (l Lote) CadastrarLote(lt Lote){

	l.IdentificadorLote = lt.IdentificadorLote
	l.Prateleira = lt.Prateleira
	l.Corredor = lt.Corredor
	l.Quantidade = lt.Quantidade
	l.DataDeEntrada = lt.DataDeEntrada
	l.DataDeValidade = lt.DataDeValidade
	l.Produto = lt.Produto

	s :=funcoes.ConverteStructEmString(l)
	r :=funcoes.ConverteDadosEmJson(s)
	repositorioDados.SalavaDadosNoRepositorio(nomeArquivo, r)
}
