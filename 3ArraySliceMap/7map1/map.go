package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[12345678] = "maria"
	aprovados[123765890] = "pedro"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[12345678])
	delete(aprovados, 123765890)
	fmt.Println(aprovados[123765890])
}
