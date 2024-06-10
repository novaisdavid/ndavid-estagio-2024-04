package repositorioDados

import (
	"os"
	"io/ioutil"
	"fmt")
	

var ficheiroDeDados *os.File

func SalavaDadosNoRepositorio(nomeRepositorio string, dadosJson []byte){
	  
	ficheiroDeDados, _ := os.OpenFile(nomeRepositorio +".csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer ficheiroDeDados.Close()

	fmt.Fprintln(ficheiroDeDados, string(dadosJson))

	ficheiroDeDados.Sync()
}


func LeDadosDoRepositorio(nomeRepositorio string) string{

	dados, erro := ioutil.ReadFile(nomeRepositorio +".csv")

	if erro != nil {
		fmt.Println("Erro ao ler o repositório de dados:", erro)
	}
	
	return string(dados)
}