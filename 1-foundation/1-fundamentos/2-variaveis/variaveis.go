package main

func main() {
	// Sintaxe básica de declaração de variáveis:
	// var nomeDaVariavel tipoDaVariavel
	// Exemplos:
	var idade int
	var nome string
	var salario float64
	var adulto bool

	// Sintaxe básica de atribuição de variáveis
	// nomeDaVariavel = valor
	// Exemplos:
	idade = 25
	nome = "João"
	salario = 2500.50
	adulto = true
	println(idade, nome, salario, adulto)

	// Sintaxe de declaração e atribuição em uma linha:
	// var nomeDaVariavel tipoDaVariavel = valor
	// Exemplo:
	var sobrenome string = "Mai"
	println(sobrenome)

	// Sintaxe de atribuição curta:
	// nomeDaVariavel := valor
	// Exemplo:
	quantidade := 10
	valor := 80.00
	println(valor, quantidade)

	// Ao declarar variáveis sem atribuir valores, cada tipo possui um valor padrão:
	// Inteiro: 0
	// Float64: 0.0
	// Booleano: false
	// String: ""
	// Interface, slice, channel, map, ponteiro, função: nil

	// Tipos Compostos:
	// Slice, map, channel: nil (vazio)
	// Array: elementos com valores padrão do tipo (ex: array de `int` com elementos `0`)
	// Struct: campos com valores padrão do tipo (ex: campos `int` com valor `0` e campos `string` com valor `""`)
}
