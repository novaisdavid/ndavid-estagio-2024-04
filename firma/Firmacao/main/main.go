package main

import (
	//"flag"
	"fmt"
	"github.com/spf13/cobra"
	servico "Firma/servico"
)

func main() {
	/*f := flag.String("f", "", "Para Matricular um formando")
	nome := flag.String("nome", "", "o nome do formando")
	emial := flag.String("email", "", "o email do formando")
	telefone := flag.Int("telefone", 0, "o telefone do formando")
	curso := flag.String("curso", "sem", "o curso do formando")
	regime := flag.String("regime", "sem", "o regime do formando (online / presencial)")
	flag.Parse()


	fmt.Println("A função chamada foi : ",*f)
	fmt.Println("NOME: ",*nome)
	fmt.Println("EMAIL: ",*emial)
	fmt.Println("TELEFONE: ",*telefone)
	fmt.Println("CURSO: ",*curso)
	fmt.Println("REGIME: ",*regime)*/

	var rootCmd = &cobra.Command{Use: "GESTFIRMA"}
	var nome, email, telefone, curso, regime string
	var cmd = &cobra.Command{
		Use:   "matricular",
		Short: " Matricular um formando",
		Run: func(cmd *cobra.Command, args []string) {
			if nome == "" {
				fmt.Println("O NOME NAO PODE SER VAZIO: ")
				return
			}
			if email == "" {
				fmt.Println("O EMAIL NÃO PODE SER VAZIO: ")
				return
			}
			if telefone == "" {
				fmt.Println("O TELEFONE NÃO PODE SER VAZIO: ")
				return
			}
			fmt.Println("O NOME : ", nome)
			fmt.Println("O EMAIL : ", email)
			fmt.Println("O TELEFONE : ", telefone)
			fmt.Println("O CURSO: ", curso)
			fmt.Println("O REGIME : ", regime)

			servico.FazerMatriculadeFormando(nome, email, telefone,curso,regime)

		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "nome do formando")
	cmd.Flags().StringVarP(&email, "email", "e", "", "email do formando")
	cmd.Flags().StringVarP(&telefone, "telefone", "t", "", "telefone do formando")
	cmd.Flags().StringVarP(&curso, "curso", "c", "", "curso do formando")
	cmd.Flags().StringVarP(&regime, "regime", "r", "", "regime do formando")
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
