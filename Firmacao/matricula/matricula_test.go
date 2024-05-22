package matricula_test

import (
	formando "Firma/formando"
	//matricula "Firma/matricula"
	//curso    "Firma/curso"
	"testing"
)

func TestCriarMatricula(t *testing.T) {
	t.Run("Verifica se recebe dados", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() == "" {
			t.Fail()
		}
	})
	t.Run("Verifica se o nome recebido esta certo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() != "Zeca" {
			t.Fail()
		}
	})
	/*t.Run("Criar Matricula do Formando", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("0001", "Ingles", "1-verbo to be 2-verbo to have", 20, "presencial")

		m.MatricularFormando(f.GetIdFormando(), c.GetIdCurso())
		//Act
        r := m.MostraEstudaMatriculado(f.GetIdFormando())

		if r.IdFormando =="" {
			t.Fail()
		}
		
		//Assert
	})*/
}
// fazer um save dee
     