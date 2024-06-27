package conversor

import "fmt"

func Conversor(numeroEmRomano string) int {
	if numeroEmRomano == "I" {
		fmt.Println("zzzzzzzz")
		return 1
	}
	if numeroEmRomano == "II" {
		fmt.Println("AAAAAAA")
		return 2
	}
	fmt.Println("OOOOOOOO")
	return -1
}
