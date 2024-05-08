package lote

import (
	"time"
	"fmt"
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

	for _, lote := range lotes {
		r:= l.ordenaDatasDoLote(lote.DataDeValidade, dataAtual)
		if r !="" {
			lot = append(lot, lote)
		}

		
	}
	fmt.Println("SAIDA dos lotes: ", lot)

	return lot
}

func (l Lote) MostraLotePorLocalizacao(lotes []Lote, lc string) Lote{

	for _, lote := range lotes {
		if lote.Localizacao == lc {
		    return lote
		}
	}

	return Lote{}

}

// fazer reornrar varios lotes por localização
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
		return data1.String()
	}

	return ""
}

/* acrescer as ideias

func dataMaisProxima(data1, data2 time.Time) time.Time {
    hoje := time.Now()
    diferenca1 := diferencaDeDatas(data1, hoje)
    diferenca2 := diferencaDeDatas(data2, hoje)
    if diferenca1 < diferenca2 {
        return data1
    }
    return data2
}
*/