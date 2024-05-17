package funcoes

import "time"

func ComparaSeDiasMenorOuIgual15(data string) bool {
	dataAtual := time.Now()
	tempoValidadeDias := 15

	data1, _ := time.Parse("2006-01-02", data)
	diferencaDatas := data1.Sub(dataAtual).Hours() / 24

	return diferencaDatas <= float64(tempoValidadeDias)

}
