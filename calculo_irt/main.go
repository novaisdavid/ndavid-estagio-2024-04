package main

import (
	clx "github.com/calculo_irt/cli"
	dadosIrt "github.com/calculo_irt/entidade"
)

func main() {
	
	tabela := []dadosIrt.TabelaIrt{
		{ValorInicial: 0, Limite: 100000, ParcelaFixa: 0, Excesso: 0, Taxa: 0},
		{ValorInicial: 1001000, Limite: 150000, ParcelaFixa: 0, Excesso: 100001, Taxa: 0.13},
		{ValorInicial: 150001, Limite: 200000, ParcelaFixa: 12500, Excesso: 150001, Taxa: 0.16},
		{ValorInicial: 200001, Limite: 300000, ParcelaFixa: 31250, Excesso: 200001, Taxa: 0.18},
		{ValorInicial: 300001, Limite: 500000, ParcelaFixa: 49250, Excesso: 300001, Taxa: 0.19},
		{ValorInicial: 500001, Limite: 1000000, ParcelaFixa: 87250, Excesso: 500001, Taxa: 0.2},
		{ValorInicial: 1000001, Limite: 1500000, ParcelaFixa: 187249, Excesso: 1000001, Taxa: 0.21},
		{ValorInicial: 1500001, Limite: 2000000, ParcelaFixa: 292249, Excesso: 1500001, Taxa: 0.22},
		{ValorInicial: 2000001, Limite: 2500000, ParcelaFixa: 402249, Excesso: 2000001, Taxa: 0.23},
		{ValorInicial: 2500001, Limite: 5000000, ParcelaFixa: 517249, Excesso: 2500001, Taxa: 0.24},
		{ValorInicial: 5000001, Limite: 10000000, ParcelaFixa: 1117249, Excesso: 5000001, Taxa: 0.245},
		{ValorInicial: 10000001, Limite: 0, ParcelaFixa: 2342248, Excesso: 10000001, Taxa: 0.25},
	}

	calculoIrt := dadosIrt.CalculoIrt{}.NewIrt(tabela)
	dadosDoIrt := dadosIrt.DadosIrt{}

	clx.Cli(calculoIrt, dadosDoIrt)

}
