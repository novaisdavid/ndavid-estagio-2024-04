package matricula

import (
	"Firma/domain/curso"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Matricula struct {
	idFormando string
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

func (m *Matricula) GetDataInicioCurso() string {
	return m.curso.GetDataInicio()
}

func (m *Matricula) GetNomeCurso() string {
	return m.curso.GetNome()
}

func (m *Matricula) GetDataFim() string {
	return m.curso.GetDataFim()
}

func (m Matricula) MostraUmEstudaMatriculado(idFormando string) Matricula {
	dados := m.LerDados()
	fm := m.ConverteEmStruct(dados)
	var structAuxiliar Matricula

	if len(fm) == 0 {
		return Matricula{}
	}

	for _, f := range fm {

		if f.idFormando == idFormando && f.GetDataFim() != "" {
			structAuxiliar = f
		}
	}

	return structAuxiliar
}

func (m Matricula) MostraTodasMatriculasComDataInicio() []Matricula {
	dados := m.LerDados()
	matriculados := m.ConverteEmStruct(dados)
	var matriculasComDataInicio []Matricula

	for _, mt := range matriculados {

		if mt.curso.GetNome() != "" && mt.curso.GetDataInicio() != "" {
			matriculasComDataInicio = append(matriculasComDataInicio, mt)
		}
	}
	return matriculasComDataInicio
}

func (m Matricula) Salvar() {

	if m.idFormando == "" || m.curso.GetIdCurso() == "" {
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

	dados := fmt.Sprintf("IdFormando: %s\nIdCurso: %s\nNome do Curso: %s\nHoras: %d\nConteudo Programatico: %s\nRegime: %s\nData Inicio: %s\nData Fim: %s\n", m.idFormando, m.curso.GetIdCurso(),
		m.curso.GetNome(), m.curso.GetHora(), m.curso.GetConteudoProgramatico(), m.curso.GetRegime(), m.curso.GetDataInicio(), m.curso.GetDataFim())
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

func (m Matricula) ConverteEmStruct(dados string) []Matricula {
	var idCurso, idFormando, dataFim, dataInicio, ctp, regime, nome string
	var horas int
	var matriculas []Matricula

	linhas := strings.Split(string(dados), "\n")
	cursosAdicionados := make(map[string]bool)

	for _, linha := range linhas {
		campos := strings.SplitN(linha, ": ", 2)
		if len(campos) == 2 {
			chave := strings.TrimSpace(campos[0])
			valor := strings.TrimSpace(campos[1])
			switch chave {
			case "IdFormando":
				idFormando = valor
			case "IdCurso":
				idCurso = valor
			case "Nome do Curso":
				nome = valor
			case "Horas":
				horas, _ = strconv.Atoi(valor)
			case "Conteudo Programatico":
				ctp = valor
			case "Regime":
				regime = valor
			case "Data Inicio":
				dataInicio = valor
			case "Data Fim":
				dataFim = valor
			}
		}

		if idFormando != "" && idCurso != "" && nome != "" && horas > 0 && ctp != "" && regime != "" && dataInicio != "" && dataFim != "" {
			if !cursosAdicionados[idCurso] {

				cur := curso.Curso{}
				cur.New(idCurso, nome, ctp, horas, regime)
				cur.DefinirDataDeInciodoCurso(&dataInicio)
				cur.DefinirDataDeFimdoCurso(&dataFim)
				matricula := Matricula{
					idFormando: idFormando,
					curso:      cur,
				}
				matriculas = append(matriculas, matricula)
				cursosAdicionados[idCurso] = true

				idFormando = ""
				idCurso = ""
				nome = ""
				horas = 0
				ctp = ""
				regime = ""
				dataInicio = ""
				dataFim = ""
			}
		}

	}

	return matriculas

}
