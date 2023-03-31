//Peça ao usuário para digitar uma string e um caractere e retorne o número de
//ocorrências desse caractere na string.

package main

import "fmt"

func main() {
	var str1, str2 string
	fmt.Println("Informe uma string e um caractere: ")
	fmt.Scan(&str1)
	fmt.Scan(&str2)
	z := 0
	for i := 0; i < len(str1); i++ {
		if string(str1[i]) == str2 {
			z++
		}
	}
	fmt.Print(z)
}
