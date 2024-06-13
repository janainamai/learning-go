package main

import "fmt"

// exemplo de método genérico que espera receber um dos tipos: int64, float64, int32
func SomaGenerica[T int64 | float64 | int32](m map[string]T) T {
	var soma T
	for _, valor := range m {
		soma += valor
	}

	return soma
}

// também posso criar uma estrutura genérica
type Numero interface {
	float64 | int64
}

// método que aceita estrutura genérica Numero
func SomaEstruturaGenerica[T Numero](numero1 T, numero2 T) T {
	return numero1 + numero2
}

// comparação genérica - comparable
// se queremos comparar igualdade de dois tipos genéricos, devemos usar o comparable
// comparable indica que os tipos são tipos comparáveis
func ComparacaoGenerica[T comparable](numero1 T, numero2 T) bool {
	return numero1 == numero2
}

func main() {
	// testando função genérica com int32
	mapInt32 := make(map[string]int32)
	mapInt32["Janaina"] = 24
	mapInt32["Heloísa"] = 23
	mapInt32["Perseu"] = 1
	mapInt32["Atena"] = 1
	mapInt32["Nix"] = 1

	resultInt32 := SomaGenerica[int32](mapInt32)
	fmt.Println("Resultado da soma int32:", resultInt32)

	// testando função genérica com estrutura genérica
	resultInt64 := SomaEstruturaGenerica[int64](24, 23)
	fmt.Println("Resultado da soma int64:", resultInt64)

	// testando comparable
	resultBool := ComparacaoGenerica[int32](2, 2)
	fmt.Println("Resultado da comparação int32", resultBool)
}
