//Peça ao usuário para digitar duas strings e retorne a concatenação das duas.

package main

import "fmt"

func main() {
	var (
		string1, string2 string
	)
	fmt.Println("Informe duas strings: ")
	fmt.Scan(&string1)
	fmt.Scan(&string2)
	s := string1 + " " + string2
	fmt.Print(s)
}
