package matricula_test

import (
	curso "Firma/curso"
	formando "Firma/formando"
	matricula "Firma/matricula"
	"testing"
)

func TestCriarMatriculaFormando(t *testing.T) {
	t.Run("Falha se formando nao recebe ID", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdFormando() == "" {
			t.Fail()
		}
	})
	t.Run("Falha se o curso nao tem ID", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdCurso() == "" {
			t.Fail()
		}
	})
	t.Run("Verifica se o ID do formando esta certo", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdFormando() != "777" {
			t.Fail()
		}
	})
	t.Run("Verifica se o ID do curso esta certo", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdCurso() != "099" {
			t.Fail()
		}
	})
	t.Run("Falha se o ID do formando estiver certo", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdFormando() == "007" {
			t.Fail()
		}
	})
	t.Run("Falha se o ID do curso estiver certo", func(t *testing.T) {
		//Arrange
		m := matricula.Matricula{}
		//Act
		m.New("777", "099")
		//Assert
		if m.GetIdCurso() == "069" {
			t.Fail()
		}
	})
	t.Run("Verifica se a Matricula do Formando foi feita", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("003", "SAD", "1-Tomada de Decisao", 40, "online", "espera")

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
		c.New("0002", "Programaçao", "1- Fluxo de dados 2-estrutura de repetição", 20, "online", "espera")

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
		c.New("", "", "", 0, "", "")

		m.New(f.GetIdFormando(), c.GetIdCurso())

		//Act
		m.Salvar()

		//Assert

		if m.GetIdFormando() != "" {
			t.Fail()
		}
	})
}
