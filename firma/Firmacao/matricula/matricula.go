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
	m.IdFormando = idFormando
	m.Curso = c
	m.Regime = r
	m.salvar()
}

func (m Matricula) MostraEstudaMatriculado(idFormando string) Matricula {
	if m.IdFormando == idFormando {
		return m
	}

	return Matricula{}
}

func (m Matricula) salvar() {
	
	dados := fmt.Sprintf("IdFormando: %s\nCurso: %s\nRegime: %s\n", m.IdFormando, m.Curso, m.Regime)
	nomeArquivo := "MatriculaFormando.txt"

	arquivo, err := os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer arquivo.Close()

	_, err = arquivo.WriteString(dados+"\n")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Dados salvos com sucesso!")
}


func (m Matricula) LerDados() string{
	nomeArquivo := "MatriculaFormando.txt"

	conteudo, erro := ioutil.ReadFile(nomeArquivo)
	if erro != nil {
		fmt.Println("Erro ao ler o arquivo:", erro)
		return ""
	}

	fmt.Println("Conteúdo do arquivo:")
	fmt.Println(string(conteudo))
	return string(conteudo)
}
