package produto_stock_test

import (
	
	produto_stock "Stock_Acme"
	"testing"
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

func teste_cadastro_produto(t *testing.T){
	//Arrange
	z := produto_stock.Produto{}
	d := "arroz"
	ca:= "local fresco"
	cat:= ""

	i := z.CadastroProduto(d,ca,cat)
	//testando(t, 1, s)
	testando2(t, 1, i)

}