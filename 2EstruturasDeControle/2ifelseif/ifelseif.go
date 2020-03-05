package main

import "fmt"

var n float64 = 0

func main() {
	menu()
}

func menu() {
	print("Digite um valor: ")
	fmt.Scan(&n)
	if n == 11 {
		println("Saiu do programa")
	} else {
		println(notaParaConceito(n))
		menu()
	}
}

func notaParaConceito(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8 {
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}
