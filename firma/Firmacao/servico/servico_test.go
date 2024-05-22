package servico_test

import (
	"Firma/formando"
	"Firma/funcoes"
	"Firma/matricula"
	"Firma/servico"
	"fmt"
	"testing"
)

func TestServicos(t *testing.T) {

	t.Run("Enviar Lembre de inicio do curso", func(t *testing.T) {
		// arrange
		matriculado := matricula.Matricula{}
		formando := formando.Formando{}
		enviou := 0

		// Act
		mt := matriculado.MostraTodasMatriculasComDataInicio()
		for _, m := range mt {
			form := formando.BuscaFormandoPorId(m.GetIdFormando())
			diasMaior := funcoes.ComparaSeDiaMaiorOuIgual2(m.GetDataInicioCurso())

			if diasMaior && m.GetIdFormando() != "" {
				servico.EnviarLembretePeloWhatsapp(m, form)
				enviou = enviou + 1

			}

		}

		// assert
		if enviou <1 {
			fmt.Println(enviou)
			t.Fail()
		}
	})
	t.Run("Verifica se curso foi avaliado", func(t *testing.T) {
		//Arrange
		//m := matricula.Matricula{}
		avaliacaoCurso := servico.AvaliarCurso("57", "90", 100)
		//Assert
		if avaliacaoCurso[0] == "" {
			fmt.Println("AVALIACAO DO CURSO NAO FOI FEITA : ", avaliacaoCurso)
			t.Fail()
		}
	})
	t.Run("Verifica se a avaliacao recebida esta correcta", func(t *testing.T) {
		//Arrange
		avaliacaoEsperada := [2]string{"80", "0"}
		//m := matricula.Matricula{}
		avaliacaoCurso := servico.AvaliarCurso("57", "90", 80)
		//Assert
		if avaliacaoCurso[0] != avaliacaoEsperada[0] {
			fmt.Println("RESULTADO DA AVALIACAO ERRADO : ", avaliacaoCurso, "ESPERADO : ", avaliacaoEsperada)
			t.Fail()
		}
	})
	t.Run("Falha se a avaliacao recebida estiver correcta", func(t *testing.T) {
		//Arrange
		avaliacaoEsperada := [2]string{"80", "0"}
		//m := matricula.Matricula{}
		avaliacaoCurso := servico.AvaliarCurso("57", "90", 50)
		//Assert
		if avaliacaoCurso[0] == avaliacaoEsperada[0] {
			fmt.Println("RESULTADO EXACTO DA AVALIACAO: ", avaliacaoCurso, "NAO ESPERADO", avaliacaoEsperada)
			t.Fail()
		}
	})
	t.Run("Verifica se a avaliacao esta acima dos 60%", func(t *testing.T) {
		//Arrange
		avaliacaoEsperada := [2]string{"80", ""}
		//m := matricula.Matricula{}
		avaliacaoCurso := servico.AvaliarCurso("57", "90", 80)
		//Assert
		if avaliacaoCurso[0] != avaliacaoEsperada[0] || avaliacaoCurso[1] != avaliacaoEsperada[1] {
			fmt.Println("AVALIACAO DE : ", avaliacaoCurso[0], "%", " AVALIACAO ESPERADA ", avaliacaoEsperada[0], "%")
			t.Fail()
		}
	})
	t.Run("Verifica se penalizacao bem aplicada para avaliacao abaixo dos 60%", func(t *testing.T) {
		//Arrange
		avaliacaoEsperada := [2]string{"50", "20"}
		//m := matricula.Matricula{}
		avaliacaoCurso := servico.AvaliarCurso("57", "90", 50)
		//Assert
		if avaliacaoCurso[0] != avaliacaoEsperada[0] || avaliacaoCurso[1] != avaliacaoEsperada[1] {
			fmt.Println("AVALIACAO DE : ", avaliacaoCurso[0], "%", " AVALIACAO ESPERADA ", avaliacaoEsperada[0], "%", " PENALIZACAO DE: ", avaliacaoCurso[1], "%", " ESPERADA ", avaliacaoEsperada[1], "%")
			t.Fail()
		}
	})
}
