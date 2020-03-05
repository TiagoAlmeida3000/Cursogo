package main

import (
	"math/rand"
	"time"
)

func main() {
	if i := numeroAleatorio(); i > 5 {
		println("Ganhou")
	} else {
		println("Perdeu")
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
