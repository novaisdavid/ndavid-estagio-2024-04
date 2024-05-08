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

// fazer reornrar varios lotes por localização
func (l *Lote) ValidadeMaisProxima() bool {
	// Regra de cliente

	// Comparar as datas aqui

	// Hoje

	// Intervalo (meses, semanas, anos, segundos)

	// l.DataDeValidade

	return true
}

/*func (l Lote) retornaDataDoLoteFormatado(data, dataAtual string) string {

	data1, _ := time.Parse("2006-01-02", data)
	data2, _ := time.Parse("2006-01-02", dataAtual)

	if data1.After(data2) {
		return data1.String()
	}

	return ""
}*/

func (l Lote) compara(data string, tempoValidadeDias int) string {
	dataAtual := time.Now()

	data1, _ := time.Parse("2006-01-02", data)
	diferencaDatas := data1.Sub(dataAtual).Hours() / 24
	if diferencaDatas <= float64(tempoValidadeDias) {
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
