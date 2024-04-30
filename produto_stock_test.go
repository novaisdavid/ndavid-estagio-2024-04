package produto_stock_test

import (
	produto_stock "Stock_Acme"
	"testing"
)

func testando(t *testing.T, esperado, actual string){
	if esperado != actual {
		t.Logf("%s != %s", esperado, actual)
		t.Fail()
	}
}

func teste_cadastro_produto(t *testing.T){
	//Arrange
	d := "Sal_hymalaiano"
	ca := "Ambiente_Frio"
	c := "Alimento"
	
	//Act
	produto_stock.(d,ca,c)
	//Assert
	
}