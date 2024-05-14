package main

import (
	"flag"
	"fmt"
)

func main()  {
	f := flag.String("f", "", "Para Matricular um formando")
	nome := flag.String("nome", "", "o nome do formando")
	emial := flag.String("email", "", "o email do formando")
	telefone := flag.Int("telefone", 0, "o telefone do formando")
	curso := flag.String("curso", "sem", "o curso do formando")
	regime := flag.String("regime", "sem", "o regime do formando (online / presencial)")
	flag.Parse()

	
	fmt.Println("A função chamada foi : ",*f)
	fmt.Println("NOME: ",*nome)
	fmt.Println("EMAIL: ",*emial)
	fmt.Println("TELEFONE: ",*telefone)
	fmt.Println("CURSO: ",*curso)
	fmt.Println("REGIME: ",*regime)
}