package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"jose":    11325.45,
		"gabriel": 15456.30,
		"pedro":   1200.0,
	}

	funcsESalarios["rafael"] = 1350
	delete(funcsESalarios, "inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
