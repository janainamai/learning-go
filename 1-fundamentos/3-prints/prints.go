package main

import "fmt"

func main() {
	// Imprime na mesma linha
	fmt.Print("Mesma")
	fmt.Print(" linha")

	// Imprime em nova linhagm
	fmt.Println("\nNova linha")

	// Converte para string
	x := 4.147
	xstring := fmt.Sprint(x)
	fmt.Printf(xstring)

	// Println não imprime valores que não sejam do tipo string utilizando símbolo de soma,
	// somente quando usamos vírgula
	// fmt.Println("O valor de x é " + x)
	fmt.Println("O valor de x é ", x)

	// Imprime valor float em uma string
	fmt.Printf("O valor de x é %f", x)

	// Imprime valor float na string utilizando apenas 2 casas após a vírgula, arredondando conforme necessário
	fmt.Printf("\nO valor de x é %.2f", x)

	// Referencia variáveis com %
	a := 1
	b := 1.9999
	c := false
	d := "texto"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
