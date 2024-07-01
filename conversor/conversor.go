package conversor

import (
	"strings"
)

type Conve struct {
	tabela map[string]int
}

func (c Conve) NewConv() Conve {
	c.tabela = make(map[string]int)

	c.tabela["I"] = 1
	c.tabela["IV"] = 4
	c.tabela["v"] = 5
	c.tabela["IX"] = 9
	c.tabela["X"] = 10
	c.tabela["L"] = 50
	c.tabela["C"] = 100
	c.tabela["D"] = 500
	c.tabela["M"] = 1000

	return c
}

func (c Conve) Conversor(numeroEmRomano string) int {

	tamanho := len(numeroEmRomano)
	total := 0
	for i := 0; i < tamanho; i++ {
		valorSaida := c.tabela[strings.ToUpper(string(numeroEmRomano[i]))]

		if i+1 < tamanho && valorSaida < c.tabela[string(numeroEmRomano[i+1])] {
			total -= valorSaida
		} else {
			total += valorSaida
		}
	}

	return total
}
