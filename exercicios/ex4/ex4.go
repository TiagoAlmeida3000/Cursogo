package main

import "fmt"

func main() {
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	r := 0

	fmt.Print("Digite a nota: ")
	fmt.Scan(&n1)
	fmt.Print("Digite a nota: ")
	fmt.Scan(&n2)
	fmt.Print("Digite a nota: ")
	fmt.Scan(&n3)
	fmt.Print("Digite a nota: ")
	fmt.Scan(&n4)

	r = (n1 + n2 + n3 + n4) / 4

	fmt.Println(r)

}
