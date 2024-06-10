package formando_test

import (
	formando "Firma/formando"
	"testing"
)

func TestCriaFormando(t *testing.T) {
	t.Run("Verifica se recebe dados para o nome", func(t *testing.T) {
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
	t.Run("Falha se o nome recebido estiver certo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() == "Zequinha" {
			t.Fail()
		}
	})
	t.Run("Verifica se recebe dados para o email", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetEmailFormando() == "" {
			t.Fail()
		}
	})
	t.Run("Verifica se o email recebido esta certo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetEmailFormando() != "Zeca@gmail.com" {
			t.Fail()
		}
	})
	t.Run("Falha se o email recebido estiver certo", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetEmailFormando() == "Zequinha@gmail.com" {
			t.Fail()
		}
	})
	t.Run("Verifica se recebe dados para o nome e email", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() == "" || f.GetEmailFormando() == "" {
			t.Fail()
		}
	})
	t.Run("Verifica se o nome e o email recebido estao certos", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() != "Zeca" || f.GetEmailFormando() != "Zeca@gmail.com" {
			t.Fail()
		}
	})
	t.Run("Falha se o nome ou email recebido estiverem certos", func(t *testing.T) {
		//Arrange
		f := formando.Formando{}
		//Act
		f.New("001", "Zeca", "Zeca@gmail.com", "923456789")
		//Assert
		if f.GetNomeFormando() == "Zequinha" || f.GetEmailFormando() == "Zequinha@gmail.com" {
			t.Fail()
		}
	})
}
