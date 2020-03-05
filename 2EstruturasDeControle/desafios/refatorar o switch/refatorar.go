package main

import "fmt"

func main() {
	menu()
}

func menu() {
	var n float64
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
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}

}
