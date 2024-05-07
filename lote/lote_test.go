package lote_test

import (
	lote "Stock_Acme/lote"
	"testing"
)

func Validar(t *testing.T, valorEsperado, valorAtual string) {

	if valorEsperado != valorAtual {
		t.Logf("%s != %s", valorEsperado, valorAtual)
		t.Fail()
	}

}

func TestMostraOsLotesComDataValidadeMaisProxima__LotesComDataValidadeMaisProxima(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/04",
		},
	}
	DataActual := "2025-01-11"
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	Validar(t, l[0].DataDeValidade, DataActual)
}

func TestFalhaAoMostraOsLotesComDataValidadeMaisProxima__Falha(t *testing.T) {
	// arrange
	lot := lote.Lote{}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/04",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2029-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11/02/04",
		},
	}

	DataActual := "2025-01-12"
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	Validar(t, l[0].DataDeValidade, DataActual)
}

// fazer localiza lote