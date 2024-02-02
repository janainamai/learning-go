package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float) inferido pelo Go

	// forma reduzida de criar uma variável
	// : declara
	// = atribui
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	// declarando constantes e variáveis em blocos de construção
	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	println(a, b, c, d)

	// declarando mais de uma variável na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// declarando variáveis de tipos diferentes na mesma linha
	g, h, i := 2, false, "texto"
	fmt.Println(g, h, i)
}
