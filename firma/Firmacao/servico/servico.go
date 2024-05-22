package servico

import (
	curso "Firma/curso"
	formando "Firma/formando"
	"Firma/funcoes"
	matricula "Firma/matricula"
	"fmt"

	//"io/ioutil"
	"math/rand"
	//"net/http"
	//"net/url"
	"strconv"
	//"strings"
	"time"
)

func FazerMatriculadeFormando(dados ...interface{}) {

	f := formando.Formando{}
	c := curso.Curso{}
	m := matricula.Matricula{}
	id := "80"

	if len(dados) < 5 {
		fmt.Println("Por favor preenchas dos campos correctamente")
		return
	}

	f.New(id, dados[0].(string), dados[1].(string), dados[2].(string))
	curso := c.BuscaUmCursoPorNomeRegime(dados[3].(string), dados[4].(string))
	jaFezUmCurso := m.MostraUmEstudaMatriculado(f.GetIdFormando())
	temDesconto := funcoes.ComparaSeDiasMenorOuIgual15(jaFezUmCurso.GetDataFim())

	if curso.GetIdCurso() == "" {
		fmt.Println("")
		fmt.Println("CURSO NÂO ENCONTRADO! FORMANDO NÂO MATRICULADO")
		return
	}

	if temDesconto {

		fmt.Println("")
		fmt.Println("NOME: ", f.GetNomeFormando())
		fmt.Println("")
		fmt.Println("====== DADOS DO CURSO ======")
		fmt.Println("CURSO: ", curso.GetNome())
		fmt.Println("CARGA HORÁRIA: ", strconv.Itoa(curso.GetHora()))
		fmt.Println("REGIME: ", curso.GetRegime())

		m.New(f.GetIdFormando(), curso)
		fmt.Println("ESTUDANTE MATRICULADO COM DESCONTO")
		m.Salvar()

		return
	}

	fmt.Println("")
	fmt.Println("NOME: ", f.GetNomeFormando())
	fmt.Println("")
	fmt.Println("====== DADOS DO CURSO ======")
	fmt.Println("CURSO: ", curso.GetNome())
	fmt.Println("CARGA HORÁRIA: ", strconv.Itoa(curso.GetHora()))
	fmt.Println("REGIME: ", curso.GetRegime())

	m.New(f.GetIdFormando(), curso)

	m.Salvar()

}

func CadastrarCursos(dados ...interface{}) {
	rand.Seed(time.Now().UnixNano())
	idCurso := rand.Intn(100)
	c := curso.Curso{}

	id := strconv.Itoa(idCurso)

	if len(dados) >= 4 {

		c.New(id, dados[0].(string), dados[1].(string), dados[2].(int), dados[3].(string))
	}

	fmt.Println("")
	fmt.Println("CURSO: ", c.GetNome())
	fmt.Println("CARGA HORÁRIA: ", c.GetHora())
	fmt.Println("REGIME: ", c.GetRegime())
	fmt.Println("CONTEUDO PROGRAMÁTICO: ", c.GetConteudoProgramatico())
	c.Salvar()
	fmt.Println("")
	fmt.Println("CURSO CADASTRADO COM SUCESSO!")

}

func FuncaoBackground() {
	m := matricula.Matricula{}
	f := formando.Formando{}

	EnviarLembretePeloWhatsapp(m, f)
}

func EnviarLembretePeloWhatsapp(m matricula.Matricula, f formando.Formando) bool {

	contEnviouMensagem := 0

	mt := m.MostraTodasMatriculasComDataInicio()
	for _, m := range mt {
		form := f.BuscaFormandoPorId(m.GetIdFormando())
		diasMaior := funcoes.ComparaSeDiaMaiorOuIgual2(m.GetDataInicioCurso())

		if diasMaior && form.GetNomeFormando() != "" {
			fmt.Println("Caro Formando  ", form.GetNomeFormando(), " o seu curso de ", m.GetNomeCurso(), " vai começar daqui há dois dias")
			contEnviouMensagem += 1

		}

	}

	if contEnviouMensagem >= 1 {
		return true
	}

	return false

	/*apiurl := "https://api.alvochat.com/instance1199/messages/chat"
	data := url.Values{}
	data.Set("to", "942678046")
	data.Set("body", "WhatsApp API on alvochat.com works good")
	data.Set("priority", "")
	data.Set("preview_url", "")
	data.Set("message_id", "")

	payload := strings.NewReader(data.Encode())

	req, _ := http.NewRequest("POST", apiurl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))*/

}
func AvaliarCurso(m matricula.Matricula, idCurso, idFormando string, avaliacao int) [2]string {
	dados := m.LerDados()
	matr := m.ConverteEmStruct(dados)
	ava := [2]string{"", ""}
	for _, x := range matr {

		if x.GetIdFormando() == idFormando && x.GetIdCurso() == idCurso && avaliacao >= 0 && avaliacao <= 100 {
			ava[0] = (strconv.Itoa(avaliacao))
		}
		if ava[0] != "" && avaliacao < 60 {

			ava[1] = (strconv.Itoa((60 - avaliacao) * 2))
		}
	}
	return ava
}
