package curso_test

import (
	"Firma/curso"
	"testing"
)

func TestCurso(t *testing.T) {
	t.Run("iniciar curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("003", "cc", "fluxo", 12, "online", "iniciado")

		//assert
		if curs.GetEstado() == "" {
			t.Fail()
		}

	})

	t.Run("iniciar curso sem dados do curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("", "", "", 0, "", "iniciado")

		//assert
		if curs.GetEstado() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem estado do  curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("003", "cc", "fluxo", 12, "online", "")

		//assert
		if curs.GetEstado() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem titulo do curso", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "", "fluxo", 12, "online", "espera")

		//assert
		if curs.GetEstado() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso sem conteudo programatico", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "", "fluxo", 12, "online", "espera")

		//assert
		if curs.GetEstado() != "" {
			t.Fail()
		}

	})

	t.Run("cadatrar curso e salvar no arquivo", func(t *testing.T) {
		//arrange
		curs := curso.Curso{}

		//act
		curs.New("103", "cc", "fluxo", 12, "online", "espera")

		curs.Salvar()
		//assert
		if curs.GetEstado() == "" {

			t.Fail()
		}

	})
}
