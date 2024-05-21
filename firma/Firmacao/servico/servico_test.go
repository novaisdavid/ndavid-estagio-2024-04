package servico_test

import (
	"Firma/formando"
	"Firma/funcoes"
	"Firma/matricula"
	"Firma/servico"
	"testing"
	"fmt"
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
			if diasMaior && form.GetNomeFormando() != "" { 
				servico.EnviarLembretePeloWhatsapp(m, form)
				enviou = enviou + 1

			}

		}

		// assert
		if enviou < 1 {
			fmt.Println(enviou)
			t.Fail()
		}
	})
}
