package matricula_test

import (
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
	t.Run("Criar Matricula do Formando", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		c := matricula.Matricula{}
		c1 := "Ingles"
		c2 := "Online"
		c.MatricularFormando(f.GetIdFormando(), c1, c2)

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
		f.New("", "", "", "")
		c := matricula.Matricula{}
		c1 := "Ingles"
		c2 := "Online"
		c.MatricularFormando(f.GetIdFormando(), c1, c2)

		//Act 
		c.Salvar()

		//Assert
		fm :=c.MostraUmEstudaMatriculado("")
		
		if fm.IdFormando !="" {
			t.Fail()
		}
	})

	t.Run("Criar Matricula do Formando sem dados do curso, não salva no repositorio", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		f.New("002", "Novais", "novais@gmail.com", "991122233")
		c := matricula.Matricula{}
		c1 := ""
		c2 := ""
		c.MatricularFormando(f.GetIdFormando(), c1, c2)

		//Act 
		c.Salvar()

		//Assert
		fm :=c.MostraUmEstudaMatriculado("002")
		
		if fm.IdFormando !="" {
			t.Fail()
		}
	})
}

     