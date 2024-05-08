package encomenda

import (
	lote   "Stock_Acme/lote"
	"fmt"
	"time"
)

type Encomenda struct {

	Cliente              string
	IdentificadorProduto string
	Quantidade           int
}


func (e Encomenda) RetiraEncomenda(en Encomenda, lotes []lote.Lote) int {
	 
	r :=en.produtoComDataValidadeProxima(lotes, en.IdentificadorProduto)
	i := en.retiraUnidadeNoLote(r, en.IdentificadorProduto, en.Quantidade)
	fmt.Println("Retirou: ", i)
	fmt.Println("LOtes: ", r)
	return 1
}


// ao escolher o produto traz com a data de expiração mais proxima
func (e Encomenda) retiraUnidadeNoLote(lotes []lote.Lote, identificadorLote string, numeroUnidadeRetirar int) int {

	for _, lote := range lotes {
		if lote.IdLote == identificadorLote && lote.NumeroDeUnidades > numeroUnidadeRetirar && numeroUnidadeRetirar > 0 {
			lote.NumeroDeUnidades = lote.NumeroDeUnidades - numeroUnidadeRetirar
			return lote.NumeroDeUnidades

		}
	}
	return -9999
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

