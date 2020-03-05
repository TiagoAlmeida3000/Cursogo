package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//numero inteiro
	fmt.Println(reflect.TypeOf(2))

	//sem sinal(só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println(reflect.TypeOf(b))

	//com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println(i1)

	var i2 rune = 'a' //representa um mapeamento da tabela Uinicode (int32)
	fmt.Println(reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é", reflect.TypeOf(49.99))

	// boolean, true false
	bo := true
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Oi"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	//string com mutipla linhas
	s2 := `Ola
	meu
	nome
	é
	tiago`

	fmt.Println(len(s2))

}
