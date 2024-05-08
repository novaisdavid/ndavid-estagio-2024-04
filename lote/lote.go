package lote

import (
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

func (l Lote) RetornaLoteComDataDeValidadeMaisProxima(lotes []Lote, dataAtual string) []Lote {
	lot := []Lote{}
	datas := []string{}
//refazer  em casa
	for _, lote := range lotes {
		data := l.ordenaDatasDoLote(lote.DataDeValidade, dataAtual)
		datas = append(datas, data)

	}

	for _, data := range datas {
		for _, lote := range lotes {
			if data == l.DataDeValidade {
				lot = append(lot, lote)
			}
		}

	}

	return lot
}

func (l *Lote) ValidadeMaisProxima() bool {
	// Regra de cliente

	// Comparar as datas aqui

	// Hoje

	// Intervalo (meses, semanas, anos, segundos)

	// l.DataDeValidade

	return true
}

func (l Lote) ordenaDatasDoLote(data, dataAtual string) string {

	data1, _ := time.Parse("2006-01-02", data)
	data2, _ := time.Parse("2006-01-02", dataAtual)
	if data1.After(data2) {
		return data2.String()
	}

	return ""
}
