package produto_stock_test

import (
	
	produto_stock "Stock_Acme"
	"testing"
	"fmt"
)

func testando(t *testing.T, esperado, actual int){

	if esperado != actual {
		t.Logf("%d != %d", esperado, actual)
		t.Fail()
	}
}


func TestCadastro_produto(t *testing.T){
	//Arrange
	z := produto_stock.Produto{}

	p := []produto_stock.Produto{
		{Identificador: 1,
		Descricao:            "arroz",
		CondicaoArmazenamento: "pao" ,
		Categoria:            "queijo",},
	}
	
	 z.CadastroProduto(p)
	 fmt.Println("PAIS: ", p)
	 

}