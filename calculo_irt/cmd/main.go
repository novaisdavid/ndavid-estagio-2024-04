package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var subsidioAlimetacao, subsidioTransporte, premios, salarioBase float64
	var diasUteis, falta int

	comandoPrincipal := &cobra.Command{Use: "cirt"}
	codigoIrt := &cobra.Command{
		Use:   "cirt [diasuteis]",
		Short: "calcular o irt de um funcionário",
		Run: func(cmd *cobra.Command, args []string) {

			if salarioBase == 0 {
				fmt.Println("deve ser informado o salário base")
				return
			}
			fmt.Println("Dias Uteis: ", diasUteis)
			fmt.Println("faltas: ", falta)
			fmt.Println("Subsidio Alimentação: ", subsidioTransporte)
			fmt.Println("Subsidio Transporte: ", subsidioTransporte)
			fmt.Println("Premios: ", premios)
			fmt.Println("Salario Base: ", salarioBase)
			//servicos.ExpedirMercadoria(idProduto, quantidade)
		},
	}

	codigoIrt.Flags().IntVarP(&diasUteis, "diasuteis", "d", 0, "dias úteis de trabalho")
	codigoIrt.Flags().IntVarP(&falta, "falta", "f", 0, "faltas")
	//saber se o subsidio é mensalou diário

	codigoIrt.Flags().Float64VarP(&subsidioAlimetacao, "subalimentacao", "a", 0, "subsidio de alimentação")
	codigoIrt.Flags().Float64VarP(&subsidioTransporte, "subtransporte", "t", 0, "subsidio de transporte")
	codigoIrt.Flags().Float64VarP(&premios, "premios", "p", 0, "prémios")
	codigoIrt.Flags().Float64VarP(&salarioBase, "salariobase", "s", 0, "salário base")

	comandoPrincipal.AddCommand(codigoIrt)

	if err := comandoPrincipal.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
