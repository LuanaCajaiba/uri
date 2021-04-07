package main

import (
	"bufio"
	"fmt"
	"os"
)

var a, b, PROD int

func main() {

	//1º Tentativa

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	PROD = a * b

	if PROD > 0 {
		fmt.Scanln("PROD = ", PROD)
	} else {
		fmt.Scanln("Presentation Error")
	}

	//2º Tentativa
	//o limite de tempo foi excedido, portanto fui verificar a lentidão na leitura da entrada.
	//o fmt.Scan não usa IO com buffer
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &a)
	fmt.Fscan(in, &b)

	PROD = a * b

	fmt.Println("PROD = ", PROD)

}
