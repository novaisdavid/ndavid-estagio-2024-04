package servico

import (
	curso "Firma/curso"
	formando "Firma/formando"
	matricula "Firma/matricula"
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

	if len(dados) == 5 {
		f.New(id, dados[0].(string), dados[1].(string), dados[2].(string))
		c.New("c001", dados[3].(string), "1-introduçao ao CCTV", 54, dados[4].(string))
	}

	m.New(f.GetIdFormando(), c.GetIdCurso())

	m.Salvar()
}
