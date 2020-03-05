package main

import "fmt"

func main() {
	i := 1
	//Go não tem aritmetica de ponteiros
	var p *int = nil

	p = &i //pegando o endereço da variavel i
	*p++   // desreferenciando o ponteiro
	i++

	fmt.Println(p, *p, i, &i)
}
