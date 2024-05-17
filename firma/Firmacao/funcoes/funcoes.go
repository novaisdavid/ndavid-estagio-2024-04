package funcoes

import (
	"time"
)

func ComparaSeDiasMenorOuIgual15(data string) bool {
	dataAtual := time.Now()
	tempoValidadeDias := 15

	data1, _ := time.Parse("2006-01-02", data)
	diferencaDatas := dataAtual.Sub(data1).Hours() / 24

	if diferencaDatas <= float64(tempoValidadeDias) && diferencaDatas > 0 {
		return true
	}
	return false

}
