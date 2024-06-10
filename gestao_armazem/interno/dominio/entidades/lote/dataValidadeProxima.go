package lote

import (
	"time"
)

func DataValidadeProxima(data string) bool {
	dataAtual := time.Now()
	janelaTempo := 3 * 30 * 24 * time.Hour // Janela de tempo de 2 meses

	dataValidade, _ := time.Parse("2006-01-02", data)

	if dataValidade.After(dataAtual) && dataValidade.Before(dataAtual.Add(janelaTempo)) {
		return true
	}
	return false
}
