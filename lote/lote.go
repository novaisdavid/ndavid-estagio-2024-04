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
	lot := []Lote {}

	for _, lote := range lotes {
		l.ordenaDatasDoLote(lote.DataDeValidade, dataAtual)
		lot = append(lot, lote)
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

func (l Lote) ordenaDatasDoLote(data, dataAtual string) []string {
	saida :=[]string {}

	date1, _ := time.Parse("2006-01-02", data)
	date2, _ := time.Parse("2006-01-02", dataAtual)
	if date1.After(date2) {
		saida = append(saida, date2.String())
	}

	return saida
}
