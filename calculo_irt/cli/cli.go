package cli

import (
	"fmt"
	"os"

	calculo "github.com/calculo_irt/entidade"
	dadosIrt "github.com/calculo_irt/entidade"
	servico "github.com/calculo_irt/servico"
	"github.com/spf13/cobra"
)

func Cli(calculoIrt calculo.CalculoIrt, dadosDoIrt dadosIrt.DadosIrt) {

	var subsidioAlimetacaoMensal, subsidioTransporteMensal, premios, salarioBase float64
	var diasUteis, falta int

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

			servico.CalcularIrt(calculoIrt, dadosDoIrt.NewDadosIrt(
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
