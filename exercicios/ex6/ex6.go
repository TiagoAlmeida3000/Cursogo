package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	var a float64
	const PI = 3.14

	fmt.Print("Raio: ")
	fmt.Scan(&r)

	a = PI * math.Pow(r, 2)

	fmt.Println(a)
}
