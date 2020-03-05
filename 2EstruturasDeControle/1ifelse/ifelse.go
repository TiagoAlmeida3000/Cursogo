package main

import "fmt"

var nota float64 = 0

func main() {
	print("Digite um valor: ")
	fmt.Scan(&nota)
	imprimirResultado(nota)
}

func imprimirResultado(nota float64) {
	if nota >= 7 {
		println("Aprovado")
	} else {
		println("Reprovado")
	}
}
