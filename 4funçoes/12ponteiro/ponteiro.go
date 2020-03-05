package main

import "fmt"

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)
}

func inc1(n int) {
	n++
}

//ponteiro é representado por um *
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta
	*n++
}
