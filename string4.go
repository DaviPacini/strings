//Peça ao usuário para digitar uma string e retorne a mesma string com as letras em
//maiúsculo.

package main

import "fmt"

func main() {
	var str, result string
	fmt.Println("Informe uma string: ")
	fmt.Scan(&str)

	for i := 0; i < len(str); i++ {
		result = result + string(str[i]-32)
	}
	fmt.Print(result)
}
