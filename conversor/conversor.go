package conversor

import (
	"fmt"
	//"strconv"
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
	valorSaida := 0
	total := 0
	sepaNumeroRomano := strings.Split(numeroEmRomano, "")

	if len(sepaNumeroRomano) == 1 {
		return c.tabela[sepaNumeroRomano[0]]
	}

	for ca, x := range sepaNumeroRomano {
		fmt.Println("CHAVE: ", ca, " ----- VAlor: ", x)
		valorSaida = c.tabela[x]

		if valorSaida < c.tabela[sepaNumeroRomano[ca+1]] {
			total -= valorSaida
		} else {
			total += valorSaida
		}
	}
	fmt.Println("VALOR TOTAL: ", total)

	return total
}
