package funcoes

import (
	"fmt"
   // "os"
   produto_stock "Stock_Acme"
    "os/exec"
    "runtime"
)

var pr = produto_stock.Produto{}

func LimpaTela(){

	var comandoLimpar string

    if runtime.GOOS == "windows" {
        comandoLimpar = "cmd"
        args := []string{"/c", "cls"}
        exec.Command(comandoLimpar, args...).Run()
    } else {
        comandoLimpar = "clear"
        exec.Command(comandoLimpar).Run()
    }
}

func MostraMenu(){

	fmt.Println("===========SISTOK ACME ===========")
	fmt.Println("=========== 1- CADASTRAR PRODUTO ===========")
	fmt.Println("=========== 2- CADASTRAR LOTE===========")
	fmt.Println("=========== 3- VER PRODUTOS CADASTRADOS ===========")
	fmt.Println("=========== 4- BUSCAR LOTE POR IDENTIFICADOR ===========")
	fmt.Println("=========== 5- BUSCAR LOTE POR DATA VALIDADE ===========")
	fmt.Println("=========== 6- VER QUANTIDADE EM UM LOTE ===========")
	fmt.Println("=========== 11- SAIR  ===========")
}

func CadastrarProduto(){
	var identificador int
	var categoria, descricao, condicaoArmazenamento string

	fmt.Println("=========== CADASTRO DE PRODUTO ===========")
	fmt.Println("=========== INFORME O IDENTICADOR DO PRODUTO ===========")
	_, err := fmt.Scanln(&identificador)
	fmt.Println("=========== INFORME A CATEGORIA  DO PRODUTO ===========")
	_, err = fmt.Scanln(&categoria)
	fmt.Println("=========== INFORME A DESCRIÇÃO DO PRODUTO ===========")
	_, err = fmt.Scanln(&descricao)
	fmt.Println("=========== INFORME A CONDIÇÃO DE ARMAZANAMENTO DO PRODUTO ===========")
	_, err = fmt.Scanln(&condicaoArmazenamento)


	if err !=nil {
		fmt.Println("Erro ao ler os dados: ", err)
	}

	p := []produto_stock.Produto{
		{Identificador            :identificador,
		Descricao                :descricao,
		CondicaoArmazenamento    :condicaoArmazenamento,
		Categoria                :categoria},
	}
	pr.CadastroProduto(p)
	fmt.Println("PRODUTO CADASTRADO COM SUCESSO!")
}

func Comeco(){
	var opcao int

	comeco := 0

	for comeco == 0 {

		_, err := fmt.Scanln(&opcao)

		if err !=nil {
			fmt.Println("Erro ao ler os dados: ", err)
		}

		if opcao == 1 {
			LimpaTela()
			fmt.Println("1")
		}

		if opcao == 2 {
			LimpaTela()
			fmt.Println("2")
		}

		if opcao == 3 {
			LimpaTela()
			fmt.Println("3")
		}

		if opcao == 4 {
			LimpaTela()
			fmt.Println("4")
		}
		
		if opcao == 5 {
			LimpaTela()
			fmt.Println("5")
		}

		if opcao == 6 {
			LimpaTela()
			fmt.Println("5")
		}

		if opcao == 11 {
			LimpaTela()
			comeco = 11
			fmt.Println("SAINDO...")
		}
	}
}



