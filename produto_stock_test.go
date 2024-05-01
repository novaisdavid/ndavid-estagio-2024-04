package produto_stock_test

import (
	
	produto_stock "Stock_Acme/produto_stock"
	"testing"
	"fmt"
)

func testando(t *testing.T, esperado, actual int){

	if esperado != actual {
		t.Logf("%d != %d", esperado, actual)
		t.Fail()
	}
}

func testando2(t *testing.T, esperado int, actual int){

	if esperado != actual {
		t.Logf("%d == %d", esperado, actual)
		t.Fail()
	}
}

func TestCadastro_produto(t *testing.T){
	//Arrange
	z := produto_stock.Produto{}

	p := produto_stock.Produto{
		Descricao:            "arroz",
		CondicaoArmazenamento: "pao" ,
		Categoria:            "queijo",
	}
	 z.CadastroProduto(p)
	 fmt.Println("PAIS: ", p)

}