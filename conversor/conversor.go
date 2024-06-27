package conversor

import (
	"fmt"
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
	separador := strings.Split(numeroEmRomano, "")
	valorSaida := 0

	for _, x := range separador {
		for chave, valor := range c.tabela {
			if chave == x {
				fmt.Println("CHAVE: ", chave, " --------- VALOR: ", valor)
				valorSaida = valorSaida + valor
			}

		}

	}
	fmt.Println("VALOR SAIDA: ", valorSaida)
	return valorSaida
}
