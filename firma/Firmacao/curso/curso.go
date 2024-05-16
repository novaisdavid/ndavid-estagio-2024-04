package curso

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Curso struct {
	idCurso              string
	titulo               string
	conteudoProgramatico string
	horas                int
	regime               string
	estado               string
	dataInicio            string
	//data fim
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

func (c *Curso) GetRegime() string {
	return c.regime
}

func (c *Curso) GetNome() string {
	return c.titulo
}

func (c *Curso) GetHora() int {

	return c.horas
}

func (c *Curso) GetConteudoProgramatico() int {

	return c.horas
}

func (c *Curso) GetEstado() string {

	return c.estado
}

func (c *Curso) IniciarCurso() {
	estado := "inciado"
	dataatual := time.Now()

	dataatualEmString := dataatual.Format("2006-01-02")

	if c.idCurso != "" && c.titulo != "" || c.horas != 0 || c.regime != "" {
		c.estado = estado
		c.dataInicio = dataatualEmString
		return
	}

}

func (c Curso) BuscaUmCursoPorNome(n string) Curso {

	dados := c.LerDados()
	cursos := c.converteEmStruct(dados)

	for _, curso := range cursos {
		if curso.titulo == n {

			return curso
		}
	}

	return Curso{}
}

func (c Curso) BuscaUmCursoPorNomeRegime(n, r string) Curso {

	dados := c.LerDados()
	cursos := c.converteEmStruct(dados)

	for _, curso := range cursos {
		if curso.titulo == n && curso.regime == r {

			return curso
		}
	}

	return Curso{}
}

func (c Curso) Salvar() {

	if c.idCurso == "" || c.titulo == "" || c.horas == 0 || c.conteudoProgramatico == "" || c.regime == "" || c.estado == "" {
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

	dados := fmt.Sprintf("IdCurso: %s\nTitulo: %s\nContéudo Pragramático: %s\nHoras: %d\nRegime: %s\n Estado do Curso: %s\n",
		c.idCurso, c.titulo, c.conteudoProgramatico, c.horas, c.regime, c.estado)

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

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	nomeArquivo := "Curso.txt"
	novoDir := filepath.Join(dir, "..", "curso")
	err = os.Chdir(novoDir)
	if err != nil {
		fmt.Println("Erro ao mudar de diretório:", err)
		return ""
	}

	conteudo, erro := ioutil.ReadFile(novoDir + "/" + nomeArquivo)
	if erro != nil {
		fmt.Println("erro: ", erro)
		return ""
	}

	return string(conteudo)
}

func (c Curso) converteEmStruct(dados string) []Curso {

	blocos := strings.Split(dados, "\n\n")
	cursosMap := make(map[int]Curso)

	for _, bloco := range blocos {
		if bloco != "" {
			linhas := strings.Split(bloco, "\n")

			curso := Curso{}

			for _, linha := range linhas {
				campos := strings.SplitN(linha, ": ", 2)
				if len(campos) == 2 {
					switch campos[0] {
					case "IdCurso":
						curso.idCurso = campos[1]
					case "Titulo":
						curso.titulo = campos[1]
					case "Contéudo Pragramático":
						curso.conteudoProgramatico = campos[1]
					case "Horas":
						curso.horas, _ = strconv.Atoi(campos[1])
					case "Regime":
						curso.regime = campos[1]
					}
				}
			}

			idCurso, _ := strconv.Atoi(curso.idCurso)
			if _, ok := cursosMap[idCurso]; !ok {
				cursosMap[idCurso] = curso
			}
		}
	}

	cursos := []Curso{}
	for _, curso := range cursosMap {
		cursos = append(cursos, curso)
	}

	return cursos
}
