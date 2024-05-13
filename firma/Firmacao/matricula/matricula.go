package matricula

import (
	 //"Firma/repositorio"
	"fmt"
	"io/ioutil"	
	"os"
)

type Matricula struct {
	IdFormando string
	Curso      string
	Regime     string
}

func (m *Matricula) MatricularFormando(idFormando, c, r string) {

	if idFormando == "" || c == "" || r =="" {
		return 
	}

	m.IdFormando = idFormando
	m.Curso = c
	m.Regime = r
}

func (m Matricula) MostraEstudaMatriculado(idFormando string) Matricula {
	if m.IdFormando == idFormando {
		return m
	}

	return Matricula{}
}

func (m Matricula) Salvar() {
	
	if m.IdFormando == "" || m.Curso =="" || m.Regime =="" {
		return
	}

	dados := fmt.Sprintf("IdFormando: %s\nCurso: %s\nRegime: %s\n", m.IdFormando, m.Curso, m.Regime)
	nomeArquivo := "MatriculaFormando.txt"

	arquivo, err := os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(dados+"\n")
	if err != nil {
		return
	}
}

func (m Matricula) LerDados() string{
	nomeArquivo := "MatriculaFormando.txt"

	conteudo, erro := ioutil.ReadFile(nomeArquivo)
	if erro != nil {
		return ""
	}

	fmt.Println(string(conteudo))

	return string(conteudo)
}
