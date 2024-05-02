package lote

import (
	produto_stock "Stock_Acme"
	//funcoes "Stock_Acme/funcoes"
	//epositorioDados "Stock_Acme/repositorioDados"
	"fmt"
)

type Lote struct{

	IdentificadorLote string
	localizacao     string
	Quantidade          int
	DataDeEntrada     string
	DataDeValidade    string
	Produto           produto_stock.Produto
}


func (l *Lote) CadastrarLote(lt Lote) {

	l.IdentificadorLote = lt.IdentificadorLote
	l.localizacao = lt.localizacao
	l.Quantidade = lt.Quantidade
	l.DataDeEntrada = lt.DataDeEntrada
	l.DataDeValidade = lt.DataDeValidade
	l.Produto = lt.Produto
	
	
}

func (l Lote) BuscaLotePorIdentificador(id string){
    
	if l.IdentificadorLote == id {

		fmt.Println("===== Lote encontrado ======== ")
		fmt.Println("Identificador do Lote: ",l.IdentificadorLote)
		fmt.Println("Localização: ",l.Prateleira)
		fmt.Println("Quantidade existente: ",l.Quantidade)
		fmt.Println("Data Entrada: ",l.DataDeEntrada)
		fmt.Println("Data Validade: ",l.DataDeValidade)
		fmt.Println("Produto { ")
		fmt.Println("Categoria: ", l.Produto.Categoria)
		fmt.Println("Descrição: ", l.Produto.Descricao)
		fmt.Println("Condição Armazenamento: ", l.Produto.CondicaoArmazenamento)
		fmt.Println("}")
		return 
	}

	fmt.Println("LOte Não encontrado!", id)
}

func (l Lote) BuscaLotePorDataValidade(data string){
    
	if l.DataDeValidade == data {

		fmt.Println("===== Lote encontrado ======== ")
		fmt.Println("Identificador do Lote: ",l.IdentificadorLote)
		fmt.Println("Localização: ",l.localizacao)
		fmt.Println("Quantidade existente: ",l.Quantidade)
		fmt.Println("Data Entrada: ",l.DataDeEntrada)
		fmt.Println("Data Validade: ",l.DataDeValidade)
		fmt.Println("Produto { ")
		fmt.Println("Categoria: ", l.Produto.Categoria)
		fmt.Println("Descrição: ", l.Produto.Descricao)
		fmt.Println("Condição Armazenamento: ", l.Produto.CondicaoArmazenamento)
		fmt.Println("}")
		return 
	}
	fmt.Println("Lote Não encontrado!", data)
}

func (l Lote) VerQuantidadeExisteNumLote(id string){

	if l.IdentificadorLote == id {
		fmt.Println("===== Lote Encontrado ======== ")
		fmt.Println("Identificador do Lote: ", l.IdentificadorLote)
		fmt.Println("Localização: ", l.localizacao)
		fmt.Println("Quantidade existente: ", l.Quantidade)
		return
	}
	fmt.Println("Lote não encontrado")
}

