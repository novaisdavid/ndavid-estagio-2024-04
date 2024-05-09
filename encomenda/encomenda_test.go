package encomenda_test

import(
	encomenda "Stock_Acme/encomenda"
	lote      "Stock_Acme/lote"
	"testing"
	"fmt"
)

func Verifica(t *testing.T, valorEsperado, valorAtual int){

	if valorEsperado != valorAtual {
		t.Logf("%d == %d", valorEsperado, valorAtual)
		t.Fail()
	}
}

func VerificaSeNaoEncomendou(t *testing.T, valorEsperado  encomenda.Encomenda){

	if valorEsperado.Cliente != "" {
		t.Log("Sem Produto Para Encomendar")
		t.Fail()
	}
}

func TestEncomendaCincoUnidadeDeUmProduto__DadosEncomenda(t *testing.T){

	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
				Cliente: "Zafir",
				IdentificadorProduto: "001",
				Quantidade: 5,
	}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-07-22",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-05-01",
			NumeroDeUnidades: 20,
			Localizacao:      "91-23-14",
		},

		{IdLote: "LOTE005",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-05-12",
			NumeroDeUnidades: 20,
			Localizacao:      "91-23-14",
		},
	}

	v := encomend.RetiraEncomenda(encomendar, lotes)

	Verifica(t, v.Quantidade, 5)
}

func TestEncomendaCincoUnidadeDeUmProdutoSemProduto__Vazio(t *testing.T){

	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
				Cliente: "Zafir",
				IdentificadorProduto: "001",
				Quantidade: 5,
	}

	lotes := []lote.Lote{


	}

	v := encomend.RetiraEncomenda(encomendar, lotes)

	VerificaSeNaoEncomendou(t, v)
}

func TestEncomendaSemCliente__Vazio(t *testing.T){

	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
	}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-05-30",
			NumeroDeUnidades: 20,
			Localizacao:      "91-23-14",
		},

	}

	v := encomend.RetiraEncomenda(encomendar, lotes)

	VerificaSeNaoEncomendou(t, v)
}

func TestDescontaNumeroUnidadeDeUmProdutoNoLote__DadosEncomendaComQuantidadeRetirada(t *testing.T){

	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
				Cliente: "Zafir",
				IdentificadorProduto: "001",
				Quantidade: 15,
	}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2025-01-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-07-12",
			NumeroDeUnidades: 20,
			Localizacao:      "91-23-14",
		},

	}

	v := encomend.RetiraEncomenda(encomendar, lotes)

	Verifica(t, v.Quantidade, 15)
}

func TestRetiraUmaQuantidadeMaiorDeProdutoDoQueExistenteEmMaisDeUmLote__DadosEcomenda(t *testing.T){

	encomend := encomenda.Encomenda{}

	encomendar := encomenda.Encomenda{
				Cliente: "Zafir",
				IdentificadorProduto: "001",
				Quantidade: 50,
	}

	lotes := []lote.Lote{

		{IdLote: "LOTE001",
			IdProduto:        "001",
			DataDeProducao:   "2022-02-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 30,
			Localizacao:      "11-02-03",
		},

		{IdLote: "LOTE002",
			IdProduto:        "002",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2025-02-11",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE003",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-02",
			NumeroDeUnidades: 20,
			Localizacao:      "11-02-04",
		},

		{IdLote: "LOTE004",
			IdProduto:        "001",
			DataDeProducao:   "2022-03-12",
			DataDeValidade:   "2024-06-12",
			NumeroDeUnidades: 50,
			Localizacao:      "91-23-14",
		},

	}

	fmt.Println("============TESTE QUANTIDADE MAIOR ========= ")
	v := encomend.RetiraEncomenda(encomendar, lotes)
	fmt.Println("QUANTIDADE A RETIRAR: ",encomendar.Quantidade)
	fmt.Println("VAlor Encomenda: ",v)

	Verifica(t, v.Quantidade, 50)
}
//teste retirar uma unidade mais acima do existe num lote
