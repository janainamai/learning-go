package main

import (
	"fmt"
	"time"
)

// mais conteúdo sobre esses termos: https://janainamai.notion.site/Go-routines-2d8dbff7369343c995cc55d3db62c6df?pvs=4

// criamos esse método para chamar e entender como funcionam as goroutines
func contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
	}
}

// ajustamos a função contador para que ele aguarde 1 segundo entre cada looping do for
// para visualizar como o runtime gerenciou as 2 threads,
// e para visualizar que enquanto uma thread dormisse no contador A, ela pudesse atuar na execução do contador B
func contadorSleep(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

func main() {
	// da forma como fiz abaixo, o sistema vai trabalhar sequencialmente e o resultado será:
	// a 0
	// a 1
	// a 2
	// a 3
	// a 4
	// b 0
	// b 1
	// b 2
	// b 3
	// b 4
	// contador("a")
	// contador("b")

	// se eu fizer com que o contador A seja executado por uma goroutine,
	// somente o B aparecerá no terminal, pois a execução do contador B foi tão rápida que nem deu tempo de criar a thread A (goroutine "go")
	// go contador("a")
	// contador("b")

	// se eu inserir um time.Sleep de 1 segundo, o sistema irá aguardar 1 segundo antes de encerrar,
	// então no terminal será possível visualizar que primeiro o contador B foi executado,
	// e depois o contador A foi executado por último, demorando um pouco mais por conta da criação da thread
	// o resultado será:
	// b 0
	// b 1
	// b 2
	// b 3
	// b 4
	// a 0
	// a 1
	// a 2
	// a 3
	// a 4
	// go contador("a")
	// contador("b")
	// time.Sleep(time.Second)

	// agora vamos ver o que acontece se eu criar duas goroutines para os contadores A e B
	// tecnicamente teremos 2 threads trabalhando simultaneamente consumindo seus próprios recursos, sem compartilhar nada
	// criamos uma nova função contador para esperar 1 segundo a cada loop do for (contadorSleep), apenas para visualizar como elas trabalham ao mesmo tempo
	// se não inserirmos o sleep uma thread vai trabalhar antes da outra por conta do delay gerado entre a criação da thread do contador A e B,
	// o resultado foi:
	// a 0
	// b 0
	// b 1
	// a 1
	// b 2
	// a 2
	// b 3
	// a 3
	// b 4
	// a 4
	// go contadorSleep("a")
	// go contadorSleep("b")
	// time.Sleep(time.Second * 10)
}
