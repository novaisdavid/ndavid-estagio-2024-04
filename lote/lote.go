package lote

import (
	"fmt"
	"time"
)

type Lote struct {
	IdLote           string
	IdProduto        string
	DataDeProducao   string
	DataDeValidade   string
	NumeroDeUnidades int
	Localizacao      string
}

func (l Lote) RetornaLoteComDataDeValidadeMaisProxima(lotes []Lote, tempoValidadeDias int) []Lote {

	lot := []Lote{}

	for _, lote := range lotes {
		r := l.compara(lote.DataDeValidade, tempoValidadeDias)

		if r != "" {
			lot = append(lot, lote)
		}

	}
	fmt.Println("SAIDA dos lotes: ", lot)

	return lot
}

func (l Lote) MostraLotePorLocalizacao(lotes []Lote, lc string) Lote {

	for _, lote := range lotes {
		if lote.Localizacao == lc {
			return lote
		}
	}

	return Lote{}

}

func (l Lote) RetiraUnidadeNoLote(lotes []Lote, identificadorLote string, numeroUnidadeRetirar int) int {

	for _, lote := range lotes {
		if lote.IdLote == identificadorLote && lote.NumeroDeUnidades > numeroUnidadeRetirar {
			return lote.NumeroDeUnidades - numeroUnidadeRetirar

		}
	}
	return -9999
}

func (l Lote) compara(data string, tempoValidadeDias int) string {
	dataAtual := time.Now()

	data1, _ := time.Parse("2006-01-02", data)
	diferencaDatas := data1.Sub(dataAtual).Hours() / 24
	if diferencaDatas <= float64(tempoValidadeDias) {
		return data1.String()
	}

	return ""
}
