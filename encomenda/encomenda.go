package encomenda

import (
	lote   "Stock_Acme/lote"
	"time"
	"fmt"
)

type Encomenda struct {

	Cliente              string
	IdentificadorProduto string
	Quantidade           int
}


func (e Encomenda) RetiraEncomenda(en Encomenda, lotes []lote.Lote) Encomenda {
	 
	produtoDataValidadeproxima := en.produtoComDataValidadeProxima1(lotes, en.IdentificadorProduto)
	unidadeRestante := en.retiraUnidadeNoLote(produtoDataValidadeproxima, en.IdentificadorProduto, en.Quantidade)

	if len(produtoDataValidadeproxima) > 0 && unidadeRestante >= 0 {
		fmt.Println("PRODUTOS MAIS PROXIMOS: ",produtoDataValidadeproxima)
		return en
	}
	
	return Encomenda{}
}

func (e Encomenda) retiraUnidadeNoLote(lotes []lote.Lote, identificadorLote string, numeroUnidadeRetirar int) int {

	unidadeSubstituta :=  numeroUnidadeRetirar
	naoConseguiuRetirarUnidades := -9999

	for _, lote := range lotes {
		if lote.IdProduto == identificadorLote && lote.NumeroDeUnidades >= unidadeSubstituta && unidadeSubstituta > 0 {
			lote.NumeroDeUnidades = lote.NumeroDeUnidades - unidadeSubstituta

			return lote.NumeroDeUnidades

		}else if lote.IdProduto == identificadorLote && lote.NumeroDeUnidades <= unidadeSubstituta && unidadeSubstituta > 0 {

			unidadeSubstituta = unidadeSubstituta - lote.NumeroDeUnidades
			lote.NumeroDeUnidades -= lote.NumeroDeUnidades
		}
	}

	return naoConseguiuRetirarUnidades
}

func (e Encomenda) produtoComDataValidadeProxima(lotes []lote.Lote, idProdduto string) []lote.Lote{
	dataAtual := time.Now()
	var indicesMaisProximos []int
	dataMaisProxima := time.Time{}
	
	for i, lote := range lotes {
		if lote.IdProduto == idProdduto {
			dataValidade, _ := time.Parse("2006-01-02", lote.DataDeValidade)

			if dataValidade.After(dataAtual) && (dataMaisProxima.IsZero() || dataValidade.Before(dataMaisProxima)) {
				dataMaisProxima = dataValidade
				indicesMaisProximos = []int{i}
			} else if dataValidade.Equal(dataMaisProxima) {
				indicesMaisProximos = append(indicesMaisProximos, i)
			}
		}
	}

	var loteEncontrado []lote.Lote
	if len(indicesMaisProximos) > 0 {
		for _, indice := range indicesMaisProximos {
			
			loteEncontrado = append(loteEncontrado, lotes[indice])
		}
	} else {
		loteEncontrado = []lote.Lote{}
	}

	return loteEncontrado
}

func (e Encomenda) produtoComDataValidadeProxima1(lotes []lote.Lote, idProduto string) []lote.Lote {
	dataAtual := time.Now()
	var lotesEncontrados []lote.Lote
	janelaTempo := 4 * 30 * 24 * time.Hour // Janela de tempo de 4 meses

	for _, lote := range lotes {
		if lote.IdProduto == idProduto {
			dataValidade, _ := time.Parse("2006-01-02", lote.DataDeValidade)

			if dataValidade.After(dataAtual) && dataValidade.Before(dataAtual.Add(janelaTempo)) {
				lotesEncontrados = append(lotesEncontrados, lote)
			}
		}
	}

	return lotesEncontrados
}

