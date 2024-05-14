package matricula

import (
	 //"Firma/repositorio"
	"fmt"
	"io/ioutil"	
	"os"
	"strings"
)

type Matricula struct {
	idFormando string
	idCurso string
}

func (m *Matricula) New(idFormando, idCurso string) *Matricula {
	
	if idFormando =="" || idCurso == "" {
		return m
	}

	m.idFormando = idFormando
	m.idCurso = idCurso
	return m

}

func (m *Matricula) GetIdFormando () string{
	return m.idFormando
}

func (m *Matricula) MatricularFormando(idFormando, idCurso string) {
	m.idFormando = idFormando
	m.idCurso = idCurso

}

// fazendo save essa função é substituida
func (m Matricula) MostraEstudaMatriculado(idFormando string) Matricula {
	if m.idFormando == idFormando {
		return m
	}

	return Matricula{}
}

func (m Matricula) MostraUmEstudaMatriculado(idFormando string) Matricula {
	d := m.LerDados()
	fm := m.converteEmStruct(d)
	fmt.Println("AQUI: ", fm)
	for _, f := range fm {
		if f.idFormando == idFormando {
			fmt.Println("EUUUUUU: ", fm)
			return f
		}
	}

	return Matricula{}
}

func (m Matricula) Salvar() {
	
	if m.idFormando == "" || m.idCurso ==""{
		return
	}

	dados := fmt.Sprintf("IdFormando: %s\nIdCurso: %s\n", m.idFormando, m.idCurso)
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
                m.idFormando = valor
            case "Curso":
                m.idCurso = valor
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
