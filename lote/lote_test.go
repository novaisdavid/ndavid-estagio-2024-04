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

func VerificaSeTemElemento(t *testing.T, valorEsperado, valorAtual int) {

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

func TestMostraOsLotesComDataValidadeMaisProxima__LotesComDataValidadeMaisProxima(t *testing.T) {
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
	tempoValidadeDias := 365
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, tempoValidadeDias)

	//assert
	Validar(t, l)
}

func TestRetornaTresLotesComDataValidadeMaisProxima__3(t *testing.T) {

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

	DataActual := 1
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	VerificaSeTemElemento(t, len(l), 3)
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

//fazer teste do retirarar unidades em lotes
// fazer teste que verifica os niveis de stoks

// fazer localiza lote
