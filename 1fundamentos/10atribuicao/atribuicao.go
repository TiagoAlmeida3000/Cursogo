package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // inferência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x // Troca o valor de x e y
	fmt.Println(x, y)

}
