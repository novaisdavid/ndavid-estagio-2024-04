package lote

import (
	produto "github.com/acmllda/interno/dominio/entidades/produto"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Lote struct {
	id           string
	produto      produto.Produto
	quantidade   int
	dataValidade string
}

func New(i string, p produto.Produto, q int, d string) *Lote {
	return &Lote{
		id:           i,
		produto:      p,
		quantidade:   q,
		dataValidade: d,
	}
}

func (l Lote) Id() string {
	return l.id
}

func (l Lote) ProdutoId() string {
	return l.produto.Id()
}

func (l Lote) Quantidade() int {
	return l.quantidade
}

func (l Lote) EncontraProdutoEmUmOuMaisLotes(idproduto string) []Lote {
	Dadoslotes := l.lerDadosArquivo()
	lotes := l.converteEmStruct(Dadoslotes)
	var loteaulixiar []Lote

	for _, lote := range lotes {
		auxiliar1, _ := strconv.Atoi(lote.produto.Id())
		auxiliar2, _ := strconv.Atoi(idproduto)

		if auxiliar1 == auxiliar2 && DataValidadeProxima(lote.dataValidade) {
			loteaulixiar = append(loteaulixiar, lote)
		}

	}

	return loteaulixiar

}

func (l Lote) RetirarProdutoNoLote(lotes []Lote, quantidade int) int {

	unidadeSubstituta := quantidade

	for _, lote := range lotes {
		if lote.quantidade >= unidadeSubstituta && unidadeSubstituta > 0 {
			lote.quantidade = lote.quantidade - unidadeSubstituta
			return lote.quantidade

		} else if lote.quantidade <= unidadeSubstituta && unidadeSubstituta > 0 {
			unidadeSubstituta = unidadeSubstituta - lote.quantidade
			lote.quantidade -= lote.quantidade
		}
	}

	return -1
}

func (l Lote) lerDadosArquivo() string {

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Erro ao obter o diretório de trabalho:", err)
		return ""
	}

	defer func() {
		err = os.Chdir(dir)
		if err != nil {
			fmt.Println("Erro ao voltar ao diretório de trabalho:", err)
		}
	}()

	nomeArquivo := "lote.txt"
	novoDir := filepath.Join(dir, "interno", "infrastructura", "ficheiros")
	err = os.Chdir(novoDir)

	if err != nil {
		fmt.Println("Erro ao mudar de diretório2:", err)
		return ""
	}

	dados, err := ioutil.ReadFile(novoDir + "/" + nomeArquivo)
	if err != nil {
		fmt.Println("Erro ao ler os dados:", err)
		return ""
	}
	novoDir = ""

	return string(dados)
}

func (l Lote) converteEmStruct(dados string) []Lote {

	blocos := strings.Split(dados, "\n\n")
	loteMap := make(map[int]Lote)
	idLote := 1

	for _, bloco := range blocos {
		if bloco != "" {
			linhas := strings.Split(bloco, "\n")

			lote := Lote{}

			for _, linha := range linhas {
				campos := strings.SplitN(linha, ": ", 2)
				if len(campos) == 2 {
					switch campos[0] {
					case "Id":
						lote.id = campos[1]
					case "Produto":
						produtoCampos := strings.Split(campos[1], ", ")
						produtoId, _ := strconv.Atoi(produtoCampos[0][4:])
						produtoNome := produtoCampos[1][6:]
						dataValidadeProduto := produtoCampos[2][14:]
						lote.produto = *produto.New(strconv.Itoa(produtoId), produtoNome, strings.Replace(dataValidadeProduto, "}", "", 1))

					case "Quantidade":
						lote.quantidade, _ = strconv.Atoi(strings.Replace(campos[1], ",", "", 1))

					case "DataValidade":
						lote.dataValidade = campos[1]
					}
				}
			}

			if _, ok := loteMap[idLote]; !ok {
				loteMap[idLote] = lote
				idLote += 1
			}
		}
	}

	lotes := []Lote{}
	for _, lote := range loteMap {
		lotes = append(lotes, lote)
	}

	return lotes
}
