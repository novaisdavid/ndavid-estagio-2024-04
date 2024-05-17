package curso_test

import (
	"Firma/curso"
	"fmt"
	"testing"
)

func TestCurso(t *testing.T) {
	t.Run("iniciar curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("003", "cc", "fluxo", 12, "online")

		curs.IniciarCurso(nil)
		//assert
		if curs.GetEstado() == "" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			t.Fail()
		}

	})

	t.Run("iniciar curso sem dados do curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("", "", "", 0, "")
		curs.IniciarCurso(nil)

		//assert
		if curs.GetEstado() == "inciado" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			t.Fail()
		}

	})

	t.Run("terminar um curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		curs.New("003", "cc", "fluxo", 12, "online")

		curs.IniciarCurso(nil)
		//act
		curs.ConcluirCurso(nil)

		//assert
		if curs.GetEstado() != "concluido" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			fmt.Println(curs.GetDataFim())
			t.Fail()
		}

	})

	t.Run("terminar um curso sem dados do curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		curs.New("", "", "", 0, "")

		curs.IniciarCurso(nil)
		//act
		curs.ConcluirCurso(nil)

		//assert
		if curs.GetEstado() == "concluido" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			fmt.Println(curs.GetDataFim())
			t.Fail()
		}

	})

	t.Run("terminar um curso sem ter de iniciar o curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}
		curs.New("003", "cc", "fluxo", 12, "online")

		//act
		curs.ConcluirCurso(nil)

		//assert
		if curs.GetEstado() == "concluido" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			fmt.Println(curs.GetDataFim())
			t.Fail()
		}

	})

	t.Run("terminar um curso sem dados do curso e sem iniciar", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}
		curs.New("", "", "", 0, "")
		curs.IniciarCurso(nil)

		//act
		curs.ConcluirCurso(nil)

		//assert
		if curs.GetEstado() == "concluido" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			fmt.Println(curs.GetDataFim())
			t.Fail()
		}

	})

	t.Run("terminar um curso e salvar no ficheiro", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}
		curs.New("003", "cc", "fluxo", 12, "online")
		curs.IniciarCurso(nil)
		curs.ConcluirCurso(nil)

		//act
		curs.Salvar()

		//assert
		if curs.GetEstado() != "concluido" {
			fmt.Println(curs.GetNome())
			fmt.Println(curs.GetHora())
			fmt.Println(curs.GetConteudoProgramatico())
			fmt.Println(curs.GetEstado())
			fmt.Println(curs.GetDataInicio())
			fmt.Println(curs.GetDataFim())
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem estado do  curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("003", "cc", "fluxo", 12, "online")

		//assert
		if curs.GetEstado() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem titulo do curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "", "fluxo", 12, "online")

		//assert
		if curs.GetIdCurso() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem conteudo programatico", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "", "fluxo", 12, "online")

		//assert
		if curs.GetIdCurso() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso e salvar no arquivo", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "cc", "fluxo", 12, "online")

		curs.Salvar()
		//assert
		if curs.GetIdCurso() == "" {

			t.Fail()
		}

	})

	t.Run("ler cursos salvos no arquivo no arquivo", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		dados := curs.LerDados()

		//assert
		if dados == "" {

			t.Fail()
		}

	})
}
