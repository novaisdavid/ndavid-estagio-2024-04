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

	var titulo, conteudoProgramatico, regimeC string
	var horas int

	cmd = &cobra.Command{
		Use:   "Novocurso",
		Short: " cadastrar um novo curso",
		Run: func(cmd *cobra.Command, args []string) {
			if titulo == "" {
				fmt.Println("O TITULO DO CURSO NAO PODE SER VAZIO: ")
				return
			}
			if conteudoProgramatico == "" {
				fmt.Println("O CONTEUDO PROGRAMÁTICO NÃO PODE SER VAZIO: ")
				return
			}
			if regimeC == "" {
				fmt.Println("O O REGIME NÃO PODE SER VAZIO: ")
				return
			}

			if horas == 0 {
				fmt.Println("AS HORAS NÃO PODE SER ZERO/VAZIO: ")
				return
			}

			servico.CadastrarCursos(titulo, conteudoProgramatico, horas, regimeC)

		},
	}

	cmd.Flags().StringVarP(&titulo, "titulo", "t", "", "o titulo do curos")
	cmd.Flags().StringVarP(&conteudoProgramatico, "conteudo", "c", "", "o conteudo programático")
	cmd.Flags().StringVarP(&regimeC, "regimec", "r", "", "regime do curso")
	cmd.Flags().IntVarP(&horas, "horas", "o", 0, "horas do curso")

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

	var idCurso, idEstudante string
	var valorAvaliacao int

	cmd = &cobra.Command{
		Use:   "Avaliarcurso",
		Short: "avaliar um curso",
		Run: func(cmd *cobra.Command, args []string) {

			if idCurso == "" {
				fmt.Println("O IDENTIFICADOR DO CURSO NAO PODE SER VAZIO: ")
				return
			}

			if idEstudante == "" {
				fmt.Println("O IDENTIFICADOR DO ESTUDANTE NAO PODE SER VAZIO: ")
				return
			}

			if valorAvaliacao == 0 {
				fmt.Println("O OVALOR DA AVALIAÇÂO NÃO PODE SER ZERO/VAZIO: ")
				return
			}

			avaliou := servico.AvaliarCurso(idCurso, idEstudante, valorAvaliacao)

			if len(avaliou) > 0 {
				fmt.Println("CURSO VALIADO!")
			}

		},
	}

	cmd.Flags().StringVarP(&idCurso, "idcurso", "i", "", "o nome do curso")
	cmd.Flags().StringVarP(&idEstudante, "idFormando", "e", "", "o nome do curso")
	cmd.Flags().IntVarP(&valorAvaliacao, "avaliacao", "n", 0, "o nome do curso")

	fmt.Println("========== LEMBRETE ENVIADO ===========")
	servico.FuncaoBackground()

	fmt.Println("============***********=============")
	fmt.Println("")
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
