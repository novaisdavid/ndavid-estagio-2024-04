package matricula_test

import (
	curso "Firma/curso"
	formando "Firma/formando"
	matricula "Firma/matricula"
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
	t.Run("Verifica se a Matricula do Formando foi feita", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("003", "SAD", "1-Tomada de Decisao", 40, "online")

		m.New(f.GetIdFormando(), c.GetIdCurso())

		//Act
		m.Salvar()

		//assert
		if m.GetIdFormando() == "" {
			t.Fail()
		}

		//Assert
	})

	t.Run("Criar Matricula do Formando sem dados do formando, não salva no repositorio", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("", "", "", "")
		c.New("0002", "Programaçao", "1- Fluxo de dados 2-estrutura de repetição", 20, "online")

		m.New(f.GetIdFormando(), c.GetIdCurso())

		//Act
		m.Salvar()

		//Assert

		if m.GetIdFormando() != "" {
			t.Fail()
		}
	})

	t.Run("Criar Matricula do Formando sem dados do curso, não salva no repositorio", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("", "", "", 0, "")

		m.New(f.GetIdFormando(), c.GetIdCurso())

		//Act
		m.Salvar()

		//Assert

		if m.GetIdFormando() != "" {
			t.Fail()
		}
	})
}
