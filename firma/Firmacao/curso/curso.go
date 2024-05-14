package curso

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Curso struct {
	idCurso            string
	titulo             string
	conteudoProgramatico string
	horas              int
	regime             string
}

func (c *Curso) New(id string, t string, cp string, h int, r string) *Curso {
	if id == "" || t == "" || h == 0 || r == "" {
		return c
	}
	c.idCurso = id
	c.titulo = t
	c.conteudoProgramatico = cp
	c.horas = h
	c.regime = r
	return c

}

func (c *Curso) GetIdCurso() string {
	return c.idCurso
}

func (c  Curso) BuscaCursoPorNome() Curso{
 return c
}

func (c Curso) Salvar() {

	if c.idCurso== "" || c.titulo == "" || c.horas == 0 || c.conteudoProgramatico == "" || c.regime =="" {
		return 
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return
	}

	novoDir := filepath.Join(dir, "..", "curso")
	err = os.Chdir(novoDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return
	}

	dados := fmt.Sprintf("IdCurso: %s\nTitulo: %s\nContéudo Pragramático: %s\nHoras: %d\nRegime: %s\n", c.idCurso,c.titulo,c.conteudoProgramatico,c.horas,c.regime)
	nomeArquivo := novoDir + "/Curso.txt"
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

func (c Curso) LerDados() string {
	nomeArquivo := "Curso.txt"

	conteudo, erro := ioutil.ReadFile(nomeArquivo)
	if erro != nil {
		return ""
	}

	fmt.Println(string(conteudo))

	return string(conteudo)
}

func (c Curso) converteEmStruct(dados string) []Curso {
	var cursos []Curso

	linhas := strings.Split(string(dados), "\n")

	for _, linha := range linhas {

		campos := strings.SplitN(linha, ": ", 2)
		if len(campos) == 2 {
			chave := strings.TrimSpace(campos[0])
			valor := strings.TrimSpace(campos[1])
			switch chave {
			case "IdCurso":
				c.idCurso = valor
			case "Titulo":
				c.titulo = valor
			case "Contéudo Pragramático":
				c.titulo = valor
				
			}
		}

		cursos = append(cursos, c)
	}

	if len(cursos) == 0 {
		return cursos
	}

	return cursos
}