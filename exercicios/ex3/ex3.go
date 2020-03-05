package main

import "fmt"

func main() {
	n1 := 0
	n2 := 0

	fmt.Print("Digite um numero: ")
	fmt.Scan(&n1)
	fmt.Print("Digite outro numero: ")
	fmt.Scan(&n2)
	fmt.Println(n1 + n2)
}
