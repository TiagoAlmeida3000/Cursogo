package main

import "fmt"

func main() {
	c := 0
	m := 0

	fmt.Print("Quantos metro?: ")
	fmt.Scan(&m)

	c = m * 100

	fmt.Printf("%v centimetros", c)
}
