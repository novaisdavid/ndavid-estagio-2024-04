package matricula

import (
	"Firma/curso"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Matricula struct {
	idFormando string
	idCurso    string
	curso      curso.Curso
}

func (m *Matricula) New(idFormando string, c curso.Curso) *Matricula {

	if idFormando == "" || c.GetIdCurso() == "" {

		return m
	}

	m.idFormando = idFormando
	m.curso = c

	return m

}

func (m *Matricula) GetIdFormando() string {

	return m.idFormando
}

func (m *Matricula) GetIdCurso() string {
	return m.curso.GetIdCurso()
}

func (m Matricula) MostraUmEstudaMatriculado(idFormando string) Matricula {
	d := m.LerDados()
	fm := m.converteEmStruct(d)
	for _, f := range fm {
		if f.idFormando == idFormando {

			return f
		}
	}

	return Matricula{}
}

func (m Matricula) Salvar() {

	if m.idFormando == "" || m.idCurso == "" {
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return
	}

	novoDir := filepath.Join(dir, "..", "matricula")
	err = os.Chdir(novoDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return
	}

	dados := fmt.Sprintf("IdFormando: %s\nIdCurso: %s\nNome do Curso: %s\nData Inicio: %s\nData Fim: %s\n", m.idFormando, m.curso.GetIdCurso(), m.curso.GetNome(),
		m.curso.GetDataInicio(), m.curso.GetDataFim())
	nomeArquivo := novoDir + "/MatriculaFormando.txt"
	arquivo, err := os.OpenFile(nomeArquivo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}

	defer arquivo.Close()

	_, err = arquivo.WriteString(dados + "\n")
	if err != nil {
		return
	}
}

func (m Matricula) LerDados() string {
	nomeArquivo := "MatriculaFormando.txt"

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	novoDir := filepath.Join(dir, "..", "matricula")
	err = os.Chdir(novoDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return ""
	}

	conteudo, erro := ioutil.ReadFile(novoDir + "/" + nomeArquivo)
	if erro != nil {
		return ""
	}

	return string(conteudo)
}

func (m Matricula) converteEmStruct(dados string) []Matricula {
	var matriculas []Matricula

	linhas := strings.Split(string(dados), "\n")

	for _, linha := range linhas {

		campos := strings.SplitN(linha, ": ", 2)
		if len(campos) == 2 {
			chave := strings.TrimSpace(campos[0])
			valor := strings.TrimSpace(campos[1])
			switch chave {
			case "IdFormando":
				m.idFormando = valor
			case "IdCurso":
				m.idCurso = valor
			}
		}

		matriculas = append(matriculas, m)
	}

	if len(matriculas) == 0 {
		return matriculas
	}

	return matriculas
}
