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
	c.tabela["II"] = 2
	c.tabela["III"] = 3
	c.tabela["IV"] = 4
	c.tabela["V"] = 5
	c.tabela["VI"] = 6
	c.tabela["VII"] = 7
	c.tabela["VIII"] = 8
	c.tabela["IX"] = 9
	c.tabela["X"] = 10
	c.tabela["XI"] = 11
	c.tabela["XII"] = 12
	c.tabela["XIII"] = 13
	c.tabela["XIV"] = 14
	c.tabela["XV"] = 15
	c.tabela["XVI"] = 16
	c.tabela["XVII"] = 17
	c.tabela["XVIII"] = 18
	c.tabela["XIX"] = 19
	c.tabela["XX"] = 20
	c.tabela["XXX"] = 30
	c.tabela["XL"] = 40
	c.tabela["L"] = 50
	c.tabela["LX"] = 60
	c.tabela["LXX"] = 70
	c.tabela["LXXX"] = 80
	c.tabela["XC"] = 90
	c.tabela["C"] = 100
	c.tabela["D"] = 500
	c.tabela["M"] = 1000

	return c
}

func (c Conve) Conversor(numeroEmRomano string) int {
	valorSaida := 0

	for i := 0; i < len(numeroEmRomano); i++ {
		valorSaida += c.tabela[strings.ToUpper(string(numeroEmRomano[i]))]
	}

	return valorSaida
}
