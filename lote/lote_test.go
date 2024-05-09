package lote_test

import (
	lote "Stock_Acme/lote"
	"testing"
	
)

func Validar(t *testing.T, valorEsperado []lote.Lote) {

	if valorEsperado == nil {
		t.Log("Falha")
		t.Fail()
	}

}

func VerificaResultado(t *testing.T, valorEsperado, valorAtual int) {

	if valorEsperado != valorAtual {
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func VerificaSeLote(t *testing.T, valorEsperado lote.Lote) {

	if valorEsperado.IdLote == "" {
		t.Log("Falha")
		t.Fail()
	}
}

func VerificaSeLoteNaoExiste(t *testing.T, valorEsperado lote.Lote) {

	if valorEsperado.IdLote != "" {
		t.Log("Falha")
		t.Fail()

	}
}

func VerificaSeProdutoLoteNaoExiste(t *testing.T, valorEsperado lote.Lote) {

	if valorEsperado.IdProduto != "" {
		t.Log("Falha")
		t.Fail()

	}
}

func VerificaSeUnidadeMaiorQue(t *testing.T, valorEsperado, valorAtual int) {

	if valorEsperado < valorAtual {
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func VerificaSeUnidadeMenorQue(t *testing.T, valorEsperado, valorAtual int) {

	if valorEsperado > valorAtual {
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func TestVerificaOsLotesComDataValidadeMaisProximaEmDias__LotesComDataValidadeMaisProximaEmDias(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-06-12",
			DataDeValidade:   "2025-07-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-01-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/04",
		},
		{IdLote: "LOTE005",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-05-11",
			NumeroDeUnidades: 0,
			Localizacao:      "21-12-11",
		},
	}
	tempoValidadeDias := 365
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, tempoValidadeDias)

	//assert
	Validar(t, l)
}

func TestVerificaDataValidadeNegativaEmDias___Falha(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-06-12",
			DataDeValidade:   "2025-07-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-01-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/04",
		},
	}
	tempoValidadeDias := -365
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, tempoValidadeDias)

	//assert
	Validar(t, l)
}

func TestRetornaTresLotesComDataValidadeMaisProxima__3Lotes(t *testing.T) {

	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},
	}

	DataActual := 1900
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	VerificaResultado(t, len(l), 3)
}

func TestMostraLotePorLocalozacao___Lote(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	localizacao := "11-02-03"

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "1-02-04",
		},
	}

	// act
	l := lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeLote(t, l)
}

func TestNaoEncontrouLotePorLocalozacaoComListaDeLotesComElementos___(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	localizacao := "13-02-03"

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "1-02-04",
		},
	}

	// act
	l := lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeLoteNaoExiste(t, l)
}

func TestNaoEncontrouLotePorLocalozacaoComListaDeLotesVazia___(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	localizacao := "13-02-03"

	lotes := []lote.Lote{}

	// act
	l := lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeLoteNaoExiste(t, l)
}

func TestLocalizaLoteComNumeroUnidadeAcimade10__Lote(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	localizacao := "21-12-11"

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 9,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 30,
			Localizacao:      "21-12-11",
		},
	}

	// act
	l := lot.MostraLotePorLocalizacao(lotes, localizacao)

	VerificaSeUnidadeMaiorQue(t, l.NumeroDeUnidades, 10)

}

func TestNaoLocalizaLoteComNumeroUnidadeAcimade10__Lote(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	localizacao := "21-12-11"

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 9,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 7,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 9,
			Localizacao:      "21-12-11",
		},
	}

	// act
	l := lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeUnidadeMenorQue(t, l.NumeroDeUnidades, 10)

}

func TestRetiraUnidadeNoLote__NumeroUnidadesRestanteNoLote(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	numeroUnidadeRetirar := 12

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	identificadorLote := "LOTE004"
	//Act
	n := lot.RetiraUnidadeNoLote(lotes, identificadorLote, numeroUnidadeRetirar)

	VerificaResultado(t, n, 87)
}

func TestNaoRetiraNumeroUnidadeNegativaNoLote__Retorna_Menos9999(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	numeroUnidadeRetirar := -12

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	identificadorLote := "LOTE004"
	//Act
	n := lot.RetiraUnidadeNoLote(lotes, identificadorLote, numeroUnidadeRetirar)

	VerificaResultado(t, n, -9999)
}

func TestNaoRetiraNumeroUnidadeNoLoteInexistente__Retorna_Menos9999(t *testing.T) {
	// arrange
	lot := lote.Lote{}
	numeroUnidadeRetirar := 12

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	identificadorLote := "LOTE006"
	//Act
	n := lot.RetiraUnidadeNoLote(lotes, identificadorLote, numeroUnidadeRetirar)

	VerificaResultado(t, n, -9999)
}

func TestAdiconarNumeroUnidadeNoLote__NumeroDeUnidadeExistenteNoLote(t *testing.T) {

	// arrange
	lot := lote.Lote{}
	numeroUnidadeAdicionar := 10

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	identificadorLote := "LOTE003"
	//Act
	n := lot.AdicionarUnidadeNoLote(lotes, identificadorLote, numeroUnidadeAdicionar)

	//assert
	VerificaResultado(t, n, 82)
}

func TestAdiconarNumeroUnidadeNoLoteNaoExistente__Retorna_Menos9999(t *testing.T) {

	// arrange
	lot := lote.Lote{}
	numeroUnidadeAdicionar := 10

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	identificadorLote := "LOTE0022"
	//Act
	n := lot.AdicionarUnidadeNoLote(lotes, identificadorLote, numeroUnidadeAdicionar)

	//assert
	VerificaResultado(t, n, -9999)
}

func TestVerificaNivelDeStockExisteEmLotes__QuantidadeLotesComStock(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 19,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	//act
	r := lot.RetornaQuantidadeLotesComStocksDisponivel(lotes)

	//assert
	VerificaResultado(t, len(r), 4)
}

func TestRetornaTresLotesComStockExistente__3_LotesComStockDisponivel(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 0,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 10,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 72,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 99,
			Localizacao:      "21-12-11",
		},
	}

	//act
	r := lot.RetornaQuantidadeLotesComStocksDisponivel(lotes)

	//assert
	VerificaResultado(t, len(r), 3)
}

func TestRetornaQuantidadeLotesComStockVazios__QuantidadeDeLotesComStockVazio(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 0,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 0,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 0,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 0,
			Localizacao:      "21-12-11",
		},
		
	}

	//act
	r := lot.RetornaQuantidadeLotesComStocksDisponivel(lotes)

	//assert
	Validar(t, r)
}

func TestLocalizaProdutoNoLote__LoteDoProduto(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 55,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 34,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "003",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 68,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "004",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 90,
			Localizacao:      "21-12-11",
		},
	}

	identificadorProduto := "002"

	//act
	p := lot.LocalizaProdutoNoLote(lotes, identificadorProduto)

	//assert
	VerificaSeLote(t, p)
}

func TestNaoLocalizaProdutoNoLote__Vazio(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 55,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 34,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "003",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 68,
			Localizacao:      "1-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "004",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 90,
			Localizacao:      "21-12-11",
		},
	}

	identificadorProduto := "008"

	//act
	p := lot.LocalizaProdutoNoLote(lotes, identificadorProduto)

	//assert
	VerificaSeProdutoLoteNaoExiste(t, p)
}

