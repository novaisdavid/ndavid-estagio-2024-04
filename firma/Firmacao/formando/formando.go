package formando

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Formando struct {
	idFormando string
	nome       string
	email      string
	telefone   string
}

func (f *Formando) New(idFormando, nome, email, telefone string) *Formando {
	f.idFormando = idFormando
	f.nome = nome
	f.email = email
	f.telefone = telefone
	return f
}
func (f Formando) GetIdFormando() string {
	return f.idFormando
}
func (f Formando) GetNomeFormando() string {
	return f.nome
}
func (f Formando) GetEmailFormando() string {
	return f.email
}

func (f Formando) BuscaFormandoPorId(id string) Formando {

	dados := f.LerDados()
	formando := f.converteEmStruct(dados)

	for _, fm := range formando {
		if fm.idFormando == id {
			return fm
		}
	}

	return Formando{}

}

func (f Formando) LerDados() string {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	nomeArquivo := "formandos.txt"
	novoDir := filepath.Join(dir, "..", "formando")
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

func (c Formando) converteEmStruct(dados string) []Formando {

	blocos := strings.Split(dados, "\n\n")
	formandoMap := make(map[int]Formando)

	for _, bloco := range blocos {
		if bloco != "" {
			linhas := strings.Split(bloco, "\n")

			formando := Formando{}

			for _, linha := range linhas {
				campos := strings.SplitN(linha, ": ", 2)
				if len(campos) == 2 {
					switch campos[0] {
					case "IdFormando":
						formando.idFormando = campos[1]
					case "Nome":
						formando.nome = campos[1]
					case "Email":
						formando.email = campos[1]
					case "Telefone":
						formando.telefone = campos[1]
					}
				}
			}

			idFormando, _ := strconv.Atoi(formando.idFormando)
			if _, ok := formandoMap[idFormando]; !ok {
				formandoMap[idFormando] = formando
			}
		}
	}

	formandos := []Formando{}
	for _, formando := range formandoMap {
		formandos = append(formandos, formando)
	}

	return formandos
}
