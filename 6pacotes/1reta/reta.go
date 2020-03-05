package main

import (
	"fmt"
	"math"
)

// Iniciando com letra maiuscula é publico (visivel dentro e fora do pacote)
// Iniciando com letra minuscula é privado (visivel apenas dentro do pacote)

//Ponto - gerara algo publico
// ponto ou _ponto - gerara algo privado

// Ponto repreesnta uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

//Distancia é responsavel por calcular a distancia linear entre dois pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}
