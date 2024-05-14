package matricula

import (
	 //"Firma/repositorio"
	"fmt"
	"io/ioutil"	
	"os"
	"strings"
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

func (m Matricula) MostraUmEstudaMatriculado(idFormando string) Matricula {
	d := m.LerDados()
	fm := m.converteEmStruct(d)
	fmt.Println("AQUI: ", fm)
	for _, f := range fm {
		if f.IdFormando == idFormando {
			fmt.Println("EUUUUUU: ", fm)
			return f
		}
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

func (m Matricula) converteEmStruct(dados string) []Matricula {
	 var matriculas []Matricula

	linhas := strings.Split(string(dados), "\n")
	fmt.Println("LINHAS: ", linhas)
	for _, linha := range linhas {

        campos := strings.SplitN(linha, ": ", 2)
        if len(campos) == 2 {
            chave := strings.TrimSpace(campos[0])
            valor := strings.TrimSpace(campos[1])
            switch chave {
            case "IdFormando":
                m.IdFormando = valor
            case "Curso":
                m.Curso = valor
            case "Regime":
                m.Regime = valor
            }
		}
		
		matriculas = append(matriculas,m)
	}
	
 fmt.Println("SAIDA: ", matriculas)
	if len(matriculas) == 0 {
		return matriculas
	}

	return matriculas
}
