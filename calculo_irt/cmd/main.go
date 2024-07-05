package main

import (
	"fmt"
	"os"

	calculo "github.com/calculo_irt"
	dadosIrt "github.com/calculo_irt/entidade"
	servico "github.com/calculo_irt/servico"
	"github.com/spf13/cobra"
)

func main() {

	var subsidioAlimetacaoMensal, subsidioTransporteMensal, premios, salarioBase float64
	var diasUteis, falta int

	tabela := []calculo.TabelaIrt{
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
	calculoIrt := calculo.CalculoIrt{}
	dadosDoIrt := dadosIrt.DadosIrt{}
	comandoPrincipal := &cobra.Command{Use: "cirt"}
	codigoIrt := &cobra.Command{
		Use:   "cirt [diasuteis]",
		Short: "calcular o irt de um funcionário",
		Run: func(cmd *cobra.Command, args []string) {

			if salarioBase < 25000 {
				fmt.Println("O salário base não pode ser a baixo de 25.000,00 kz")
				return
			}

			if diasUteis <= 0 || diasUteis > 31 {
				fmt.Println("Os dias úteis de trabalho não pode ser negativo ou maior que 31 dias")
				return
			}

			if falta < 0 || falta > diasUteis {
				fmt.Println("O número de faltas não pode ser negativo ou maior que os dias úteis de trabalho")
				return
			}

			if subsidioAlimetacaoMensal < 0 || subsidioAlimetacaoMensal > salarioBase {
				fmt.Println("Os subsídio de alimentção não podem ser negativos ou maior que o salário base.")
				return
			}
			if subsidioTransporteMensal < 0 || subsidioTransporteMensal > salarioBase {
				fmt.Println("Os subsídio de transporte não podem ser negativos ou maior que o salário base.")
				return
			}

			servico.CalcularIrt(calculoIrt.NewIrt(tabela), dadosDoIrt.NewDadosIrt(
				diasUteis, falta,
				subsidioAlimetacaoMensal, subsidioTransporteMensal,
				premios, salarioBase))
		},
	}

	codigoIrt.Flags().IntVarP(&diasUteis, "diasuteis", "d", 22, "dias úteis de trabalho")
	codigoIrt.Flags().IntVarP(&falta, "falta", "f", 0, "faltas")
	codigoIrt.Flags().Float64VarP(&subsidioAlimetacaoMensal, "subalimentacao", "a", 0, "subsidio de alimentação")
	codigoIrt.Flags().Float64VarP(&subsidioTransporteMensal, "subtransporte", "t", 0, "subsidio de transporte")
	codigoIrt.Flags().Float64VarP(&premios, "premios", "p", 0, "prémios")
	codigoIrt.Flags().Float64VarP(&salarioBase, "salariobase", "s", 0, "salário base")

	comandoPrincipal.AddCommand(codigoIrt)

	if err := comandoPrincipal.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
