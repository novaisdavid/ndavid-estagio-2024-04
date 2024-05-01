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
		Prateleira : "P1",       
		Corredor: "C2",          
		DataDeEntrada : "29/12/2002" ,   
		DataDeValidade: "21/23/2024" ,   
		Produto: p   ,       
	}

	z.CadastrarLote(l)


	fmt.Println("p: ",p, " ..."," l: ",l)
}

func TestVerLoteExistente(t *testing.T){
	//a:=lote.VerLoteExistente()
	z:= lote.Lote{}

	z.VerLoteExistente()
	//fmt.Println(a.Corredor)
}