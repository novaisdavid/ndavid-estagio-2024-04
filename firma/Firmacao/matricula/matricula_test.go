package matricula_test

import (
	formando "Firma/formando"
	matricula "Firma/matricula"
	curso    "Firma/curso"
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
	t.Run("Criar Matricula do Formando", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c := matricula.Matricula{}
		c1 := "Ingles"
		c.MatricularFormando(f.GetIdFormando(), c1)

		//Act
		c.Salvar()

		//assert
		if c.LerDados() == "" {
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
		c.New("0002","Programaçao", "1- Fluxo de dados 2-estrutura de repetição", 20, "online")

		m.MatricularFormando(f.GetIdFormando(), c.GetIdCurso())

		//Act 
		m.Salvar()

		//Assert
		fm :=m.MostraUmEstudaMatriculado("")
		
		if fm.GetIdFormando() !="" {
			t.Fail()
		}
	})

	t.Run("Criar Matricula do Formando sem dados do curso, não salva no repositorio", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		c := curso.Curso{}
		m := matricula.Matricula{}

		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c.New("0001", "Ingles", "1-verbo to be 2-verbo to have", 20, "presencial")

		m.MatricularFormando(f.GetIdFormando(), c.GetIdCurso())

		//Act 
		m.Salvar()

		//Assert
		fm :=m.MostraUmEstudaMatriculado("002")
		
		if fm.GetIdFormando()!="" {
			t.Fail()
		}
	})
}

     