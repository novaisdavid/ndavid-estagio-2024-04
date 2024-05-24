package main

import (
	servico "Firma/servico"
	"fmt"

	"github.com/spf13/cobra"
)

func main() {

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

			servico.FazerMatriculadeFormando(nome, email, telefone, curso, regime)

		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "nome do formando")
	cmd.Flags().StringVarP(&email, "email", "e", "", "email do formando")
	cmd.Flags().StringVarP(&telefone, "telefone", "t", "", "telefone do formando")
	cmd.Flags().StringVarP(&curso, "curso", "c", "", "curso do formando")
	cmd.Flags().StringVarP(&regime, "regime", "r", "", "regime do formando")
	rootCmd.AddCommand(cmd)

	var nomeCurso string

	cmd = &cobra.Command{
		Use:   "Iniciarcurso",
		Short: "iniciar curso",
		Run: func(cmd *cobra.Command, args []string) {

			if nomeCurso == "" {
				fmt.Println("O NOME DO CURSO NAO PODE SER VAZIO: ")
				return
			}

			servico.IniciarCurso(nomeCurso)

		},
	}

	cmd.Flags().StringVarP(&nomeCurso, "nomeC", "n", "", "o nome do curso")
	rootCmd.AddCommand(cmd)

	var nomeCursoC string

	cmd = &cobra.Command{
		Use:   "Concluircurso",
		Short: "terminar curso",
		Run: func(cmd *cobra.Command, args []string) {

			if nomeCursoC == "" {
				fmt.Println("O NOME DO CURSO NAO PODE SER VAZIO: ")
				return
			}

			servico.ConcluirCursoCurso(nomeCursoC)

		},
	}

	cmd.Flags().StringVarP(&nomeCursoC, "nomeCC", "n", "", "o nome do curso")
	rootCmd.AddCommand(cmd)

	var idCurso, idFormando string
	var valorAvaliacao int

	cmd = &cobra.Command{
		Use:   "Avaliarcurso",
		Short: "avaliar um curso",
		Run: func(cmd *cobra.Command, args []string) {

			if idCurso == "" {
				fmt.Println("O IDENTIFICADOR DO CURSO NAO PODE SER VAZIO: ")
				return
			}

			if idFormando == "" {
				fmt.Println("O IDENTIFICADOR DO ESTUDANTE NAO PODE SER VAZIO: ")
				return
			}

			if valorAvaliacao == 0 {
				fmt.Println("O OVALOR DA AVALIAÇÂO NÃO PODE SER ZERO/VAZIO: ")
				return
			}

			avaliou := servico.AvaliarCurso(idCurso, idFormando, valorAvaliacao)

			if len(avaliou) > 0 {
				fmt.Println("CURSO AVALIADO!")
			}

		},
	}

	cmd.Flags().StringVarP(&idCurso, "idcurso", "i", "", "o identificador do curso")
	cmd.Flags().StringVarP(&idFormando, "idFormando", "f", "", "o identificador do formando")
	cmd.Flags().IntVarP(&valorAvaliacao, "avaliacao", "a", 0, "nota da avaliação")

	fmt.Println("========== LEMBRETE ENVIADO ===========")
	servico.FuncaoBackground()

	fmt.Println("============***********=============")
	fmt.Println("")
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
