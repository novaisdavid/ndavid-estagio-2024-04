package lote_test

import (
	
	produto_stock "Stock_Acme"
	lote "Stock_Acme/lote"
	"fmt"
	"testing"
)

func testando(t *testing.T, esperado, actual int){

	if esperado != actual {
		t.Logf("%d != %d", esperado, actual)
		t.Fail()
	}
}

func TestCadastrarLote(t *testing.T){
	//produtos := produto_stock.Produto{}
	z:= lote.Lote{}
	p := produto_stock.Produto{
		Descricao:            "arroz",
		CondicaoArmazenamento: "pao" ,
		Categoria:            "queijo",
	}
	
	l:= lote.Lote{
		IdentificadorLote: "LOTE001",
		Prateleira : "P2",       
		Corredor: "C2", 
		Quantidade: 100,         
		DataDeEntrada : "29/12/2002" ,   
		DataDeValidade: "21/23/2024" ,   
		Produto: p   ,       
	}

	z.CadastrarLote(l)


	fmt.Println("p: ",p, " ..."," l: ",l)
}

func TestVerUmLote(t *testing.T){
	

}