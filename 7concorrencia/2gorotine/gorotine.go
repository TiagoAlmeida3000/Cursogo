package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	fale("Maria", "Pq vc não fala comigo", 3)
	fale("João", "So posso falar depois com vc", 1)

	go fale("Maria", "Ei...", 500)
	go fale("Joao", "Opa...", 500)
	time.Sleep(time.Second * 5)
	fmt.Println("Fim")
}
