package lote

import (
	produto_stock "Stock_Acme"
	"fmt"
	"Stock_Acme/funcoes"
	"Stock_Acme/repositorioDados"
)

type Lote struct{

	IdentificadorLote string
	Prateleira        string
	Corredor          string
	DataDeEntrada     string
	DataDeValidade    string
	Produto           produto_stock.Produto
}

var nomeArquivo = "lote"

func (l Lote) CadastrarLote(lt Lote){

	l.IdentificadorLote = lt.IdentificadorLote
	l.Prateleira = lt.Prateleira
	l.Corredor = lt.Corredor
	l.DataDeEntrada = lt.DataDeEntrada
	l.DataDeValidade = lt.DataDeValidade
	l.Produto = lt.Produto

	s :=funcoes.ConverteStructEmString(l)
	r :=funcoes.ConverteDadosEmJson(s)
	repositorioDados.SalavaDadosNoRepositorio(nomeArquivo, r)
}

func (l Lote) VerLoteExistente(){

	r := repositorioDados.LeDadosDoRepositorio(nomeArquivo)
	 s := funcoes.ColocaDadosNoVector(r)// converter em json ou em struct
	fmt.Println("corredor em vector: ", s)
}
