package main

import "fmt"

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudavel: %t", tv50, tv32, sorvete, !sorvete)
}

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2    // E
	comprarTv32 := trab1 != trab2    // OU exclusivo
	comprarSorvete := trab1 || trab2 // OU

	return comprarTv50, comprarTv32, comprarSorvete
}
