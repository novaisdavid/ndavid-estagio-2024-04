package curso_test

import (
	"Firma/domain/curso"
	"fmt"
	"testing"
)

func TestCurso(t *testing.T) {
	t.Run("iniciar curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}
		datainicio := "2024-05-31"
		curs.New("003", "cc", "fluxo", 12, "online")
		curs.DefinirDataDeInciodoCurso(&datainicio)

		//act
		curs.IniciarCurso()
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
		curs.IniciarCurso()

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
		datinicio := "2024-04-01"
		data := "2024-05-12"
		curs.New("003", "cc", "fluxo", 12, "online")
		curs.DefinirDataDeInciodoCurso(&datinicio)
		curs.IniciarCurso()
		curs.DefinirDataDeFimdoCurso(&data)
		//act
		curs.ConcluirCurso()

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
		data := "2024-05-17"
		curs.New("", "", "", 0, "")
		curs.DefinirDataDeFimdoCurso(&data)
		curs.IniciarCurso()
		//act
		curs.ConcluirCurso()

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
		data := "2024-05-17"
		curs.New("003", "cc", "fluxo", 12, "online")
		curs.DefinirDataDeFimdoCurso(&data)

		//act
		curs.ConcluirCurso()

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
		data := "2024-05-17"
		curs.New("", "", "", 0, "")
		curs.DefinirDataDeFimdoCurso(&data)

		//act
		curs.ConcluirCurso()

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
		datafim := "2024-05-12"
		datainicio := "2024-04-06"
		curs.New("003", "cc", "fluxo", 12, "online")
		curs.DefinirDataDeInciodoCurso(&datainicio)
		curs.DefinirDataDeFimdoCurso(&datafim)
		curs.IniciarCurso()

		//act
		curs.ConcluirCurso()
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
