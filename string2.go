//Peça ao usuário para digitar uma string e retorne o número de caracteres nessa string.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Informe uma string: ")
	scanner.Scan()
	str := scanner.Text()
	fmt.Print("len(str) = ", len(str))
}
