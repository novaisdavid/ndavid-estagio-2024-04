package main

import (
	conversor "convrsorromano"
	"fmt"
)

func main() {
	numeroParaConversao := []string{"III", "IV", "IX", "LVIII", "MCMXCIV", "XXI", "XL"}
	conve := conversor.Conve{}.NewConv()

	fmt.Println("ROMANO ==========> NATURAL")
	for _, valor := range numeroParaConversao {

		fmt.Println(valor, " ==========> ", conve.Conversor(valor))
	}
}
