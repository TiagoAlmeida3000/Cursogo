package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nseparando")

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print("par ")
		} else {
			fmt.Print("impar")
		}
	}

	// for {
	// 	fmt.Println("Para sempre...")
	// 	time.Sleep(time.Second)
	// }

}
