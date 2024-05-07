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


func VerificaSeTemElemento(t *testing.T, valorEsperado, valorAtual int){

	if valorEsperado != valorAtual{
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func VerificaSeLote(t *testing.T, valorEsperado lote.Lote){

	if valorEsperado.IdLote == "" {
		t.Log("Falha")
		t.Fail()
	}
}

func VerificaSeLoteNaoExiste(t *testing.T, valorEsperado lote.Lote){
	
	if valorEsperado.IdLote != "" {
		t.Log("Falha")
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
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2019-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-01",
			NumeroDeUnidades: 20,
			Localizacao:      "12/22/01",
		},
	}
	DataActual := "2024-02-11"
	//DataActual := "2026-02-11"
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	Validar(t, l)
}

func TestNaoEncontrouOsLotesComDataValidadeMaisProxima__(t *testing.T) {
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
			Localizacao:      "1-02-04",
		},
	}

	DataActual := "2030-01-12"
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	VerificaSeTemElemento(t,len(l), 0)
}

func TestRetornaTresLotesComDataValidadeMaisProxima__3(t *testing.T){

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

	DataActual := "2024-05-06"
	//act
	l := lot.RetornaLoteComDataDeValidadeMaisProxima(lotes, DataActual)

	//assert
	VerificaSeTemElemento(t,len(l), 3)
}

// fazer localiza lote

func TestMostraLotePorLocalozacao___Lote(t *testing.T){
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
	l:= lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeLote(t,l)
}

func TestNaoEncontrouLotePorLocalozacao___(t *testing.T){
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
	l:= lot.MostraLotePorLocalizacao(lotes, localizacao)

	//assert
	VerificaSeLoteNaoExiste(t,l)
}