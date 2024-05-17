package servico

import (
	curso "Firma/curso"
	formando "Firma/formando"
	matricula "Firma/matricula"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func FazerMatriculadeFormando(dados ...interface{}) {

	rand.Seed(time.Now().UnixNano())
	idEstudante := rand.Intn(100)
	f := formando.Formando{}
	c := curso.Curso{}
	m := matricula.Matricula{}
	id := strconv.Itoa(idEstudante)

	if len(dados) < 5 {
		fmt.Println("Por favor preenchas dos campos correctamente")
		return
	}

	f.New(id, dados[0].(string), dados[1].(string), dados[2].(string))
	curso := c.BuscaUmCursoPorNomeRegime(dados[3].(string), dados[4].(string))
	if curso.GetIdCurso() == "" {
		fmt.Println("")
		fmt.Println("CURSO NÂO ENCONTRADO! FORMANDO NÂO MATRICULADO")
		return
	}

	fmt.Println("")
	fmt.Println("NOME: ", f.GetNomeFormando())
	fmt.Println("")
	fmt.Println("====== DADOS DO CURSO ======")
	fmt.Println("CURSO: ", curso.GetNome())
	fmt.Println("CARGA HORÁRIA: ", strconv.Itoa(curso.GetHora()))
	fmt.Println("REGIME: ", curso.GetRegime())

	m.New(f.GetIdFormando(), curso.GetIdCurso())

	m.Salvar()
}

func CadastrarCursos(dados ...interface{}) {
	rand.Seed(time.Now().UnixNano())
	idCurso := rand.Intn(100)
	c := curso.Curso{}

	id := strconv.Itoa(idCurso)

	if len(dados) >= 4 {

		c.New(id, dados[0].(string), dados[1].(string), dados[2].(int), dados[3].(string))
	}

	fmt.Println("")
	fmt.Println("CURSO: ", c.GetNome())
	fmt.Println("CARGA HORÁRIA: ", c.GetHora())
	fmt.Println("REGIME: ", c.GetRegime())
	c.Salvar()
	fmt.Println("")
	fmt.Println("CURSO CADASTRADO COM SUCESSO!")

}
