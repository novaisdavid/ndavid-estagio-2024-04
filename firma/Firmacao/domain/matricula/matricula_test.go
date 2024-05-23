package matricula_test

import (
	curso "Firma/domain/curso"
	formando "Firma/domain/formando"
	"Firma/funcoes"
	matricula "Firma/domain/matricula"
	"fmt"
	"testing"
)

func TestCriarMatriculaFormando(t *testing.T) {
	t.Run("Falha se formando nao recebe ID", func(t *testing.T) {
		//Arrange
		c := curso.Curso{}
		m := matricula.Matricula{}
		c.New("003", "SAD", "1-Tomada de Decisao", 40, "online")
		//Act
		m.New("777", c)
		//Assert
		if m.GetIdFormando() == "" {
			t.Fail()
		}
	})
	t.Run("Falha se o curso nao tem ID", func(t *testing.T) {
		//Arrange
		c := curso.Curso{}
		m := matricula.Matricula{}
		c.New("", "SAD", "1-Tomada de Decisao", 40, "online")
		//Act
		m.New("777", c)
		//Assert
		if m.GetIdCurso() != "" {
			t.Fail()
		}
	})
	t.Run("Verifica se o ID do formando esta certo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("87", "SAD", "1-Tomada de Decisao", 40, "online")
		//Act
		m.New(f.GetIdFormando(), c)
		//Assert
		if m.GetIdFormando() != "001" {
			t.Fail()
		}
	})
	t.Run("Verifica se o ID do curso esta certo", func(t *testing.T) {
		//Arrange
		c := curso.Curso{}
		m := matricula.Matricula{}
		c.New("87", "SAD", "1-Tomada de Decisao", 40, "online")
		//Act
		m.New("777", c)
		//Assert
		if m.GetIdCurso() != "87" {
			t.Fail()
		}
	})
	t.Run("Falha se o ID do formando estiver certo", func(t *testing.T) {
		//Arrange
		c := curso.Curso{}
		m := matricula.Matricula{}
		//Act
		m.New("777", c)
		//Assert
		if m.GetIdFormando() == "007" {
			t.Fail()
		}
	})
	t.Run("Falha se o ID do curso estiver certo", func(t *testing.T) {
		//Arrange
		c := curso.Curso{}
		m := matricula.Matricula{}
		//Act
		m.New("777", c)
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
		c.New("003", "SAD", "1-Tomada de Decisao", 40, "online")

		m.New(f.GetIdFormando(), c)

		//Act
		m.Salvar()

		//assert
		if m.GetIdFormando() == "" {
			t.Fail()
		}

	})

	t.Run("Criar Matricula do Formando sem dados do formando, não salva no repositorio", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("", "", "", "")
		c.New("0002", "Programaçao", "1- Fluxo de dados 2-estrutura de repetição", 20, "online")

		m.New(f.GetIdFormando(), c)

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

		m.New(f.GetIdFormando(), c)

		//Act
		m.Salvar()

		//Assert

		if m.GetIdFormando() != "" {
			t.Fail()
		}
	})

	t.Run("ler dados ddo arquivo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("90", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("57", "IT", "LAN", 120, "online")

		m.New(f.GetIdFormando(), c)
		m.Salvar()
		//Act
		m.LerDados()

		//Assert

		if m.MostraUmEstudaMatriculado("90") == (matricula.Matricula{}) {
			t.Fail()
		}
	})

	t.Run("matricular com desconto", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}
		f.New("90", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("34", "CCNA", "LAN, WAN", 120, "online")

		m.New(f.GetIdFormando(), c)
		e := m.MostraUmEstudaMatriculado(f.GetIdFormando())
		fmt.Println(e)
		desconto := funcoes.ComparaSeDiasMenorOuIgual15(e.GetDataFim())

		//Assert
		if !desconto {
			fmt.Println("TEM DESCONTO? ", desconto)
			t.Fail()

		}
	})
}
