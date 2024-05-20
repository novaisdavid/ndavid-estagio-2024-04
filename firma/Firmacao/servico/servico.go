package servico

import (
	curso "Firma/curso"
	formando "Firma/formando"
	"Firma/funcoes"
	matricula "Firma/matricula"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func FazerMatriculadeFormando(dados ...interface{}) {

	f := formando.Formando{}
	c := curso.Curso{}
	m := matricula.Matricula{}
	id := "90"

	if len(dados) < 5 {
		fmt.Println("Por favor preenchas dos campos correctamente")
		return
	}

	f.New(id, dados[0].(string), dados[1].(string), dados[2].(string))
	curso := c.BuscaUmCursoPorNomeRegime(dados[3].(string), dados[4].(string))
	jaFezUmCurso := m.MostraUmEstudaMatriculado(f.GetIdFormando())
	temDesconto := funcoes.ComparaSeDiasMenorOuIgual15(jaFezUmCurso.GetDataFim())

	if curso.GetIdCurso() == "" {
		fmt.Println("")
		fmt.Println("CURSO NÂO ENCONTRADO! FORMANDO NÂO MATRICULADO")
		return
	}

	if temDesconto {

		fmt.Println("")
		fmt.Println("NOME: ", f.GetNomeFormando())
		fmt.Println("")
		fmt.Println("====== DADOS DO CURSO ======")
		fmt.Println("CURSO: ", curso.GetNome())
		fmt.Println("CARGA HORÁRIA: ", strconv.Itoa(curso.GetHora()))
		fmt.Println("REGIME: ", curso.GetRegime())

		m.New(f.GetIdFormando(), curso)
		fmt.Println("ESTUDANTE MATRICULADO COM DESCONTO")
		m.Salvar()

		return
	}

	fmt.Println("")
	fmt.Println("NOME: ", f.GetNomeFormando())
	fmt.Println("")
	fmt.Println("====== DADOS DO CURSO ======")
	fmt.Println("CURSO: ", curso.GetNome())
	fmt.Println("CARGA HORÁRIA: ", strconv.Itoa(curso.GetHora()))
	fmt.Println("REGIME: ", curso.GetRegime())

	m.New(f.GetIdFormando(), curso)

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
	fmt.Println("CONTEUDO PROGRAMÁTICO: ", c.GetConteudoProgramatico())
	c.Salvar()
	fmt.Println("")
	fmt.Println("CURSO CADASTRADO COM SUCESSO!")

}
