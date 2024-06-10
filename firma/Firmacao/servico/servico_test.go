package servico_test

import (
	"Firma/formando"
	"Firma/funcoes"
	"Firma/matricula"
	"Firma/servico"
	"testing"
)

func TestServicos(t *testing.T) {

	t.Run("Enviar Lembre de inicio do curso", func(t *testing.T) {
		// arrange
		matriculado := matricula.Matricula{}
		formando := formando.Formando{}
		enviou := false

		// Act
		mt := matriculado.MostraTodasMatriculasComDataInicio()
		for _, m := range mt {
			form := formando.BuscaFormandoPorId(m.GetIdFormando())
			diasMaior := funcoes.ComparaSeDiaMaiorOuIgual2(m.GetDataInicioCurso())

			if diasMaior && form.GetNomeFormando() != "" {
				enviou = servico.EnviarLembretePeloWhatsapp(m, form)

			}

		}

		// assert
		if !enviou {
			t.Fail()
		}
	})
}
