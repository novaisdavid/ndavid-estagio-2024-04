package produto_stock_test

import (
	produto_stock "Stock_Acme"
	"testing"
)

func Validar(t *testing.T, valorEsperado produto_stock.Produto) {
	if valorEsperado.IdProduto == "" {
		t.Log("Falha")
		t.Fail()
	}
}
func ValidarSeNaoExiste(t *testing.T, valorEsperado produto_stock.Produto) {
	if valorEsperado.IdProduto != "" {
		t.Log("Falha")
		t.Fail()
	}
}

func TestVerificaSeProdutoExiste_Produto(t *testing.T) {
	//arrange
	produto1 := produto_stock.Produto{}
	produtos := []produto_stock.Produto{
		{
			IdProduto:      "001",
			NomeDoProduto:  "CocaColaEmLata330ml",
			Categoria:      "Bebidas",
			DataDeProducao: "2022-05-01",
			DataDeValidade: "2025-05-01",
		},
		{
			IdProduto:      "002",
			NomeDoProduto:  "SpriteEmLata330ml",
			Categoria:      "Bebidas",
			DataDeProducao: "2022-06-01",
			DataDeValidade: "2025-06-01",
		},
	}
	idProduto := "002"
	//act
	produtoEncontrado := produto1.VerificaSeProdutoExiste(produtos, idProduto)
	//Assert
	Validar(t, produtoEncontrado)
}

func TestVerificaSeProdutoNaoExiste_Nada(t *testing.T) {
	//arrange
	produto1 := produto_stock.Produto{}
	produtos := []produto_stock.Produto{
		{
			IdProduto:      "001",
			NomeDoProduto:  "CocaColaEmLata330ml",
			Categoria:      "Bebidas",
			DataDeProducao: "2022-05-01",
			DataDeValidade: "2025-05-01",
		},
		{
			IdProduto:      "002",
			NomeDoProduto:  "SpriteEmLata330ml",
			Categoria:      "Bebidas",
			DataDeProducao: "2022-06-01",
			DataDeValidade: "2025-06-01",
		},
	}
	idProduto := ""
	//act
	produtoEncontrado := produto1.VerificaSeProdutoExiste(produtos, idProduto)
	//Assert
	ValidarSeNaoExiste(t, produtoEncontrado)
}
