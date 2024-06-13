package main

import (
	"fmt"
	"time"
)

func main() {
	// criando um canal "queue" para transmitir números inteiros
	queue := make(chan int)

	// criando uma goroutine anônima
	go func() {
		i := 0

		for {
			time.Sleep(time.Second)
			// a cada segundo envia um valor "i" para o canal "queue"
			queue <- i
			// o for só vai prosseguir quando alguém acessar queue, considerando a anotação de queue abaixo #1-sqlpuro
			i++
		}

	}()

	for x := range queue {
		// no momento em que imprimimos x, o canal é liberado para receber outro valor
		fmt.Println(x)
	}

}

// #1-sqlpuro Quando um valor é enviado para o canal, ele é armazenado temporariamente no canal até ser lido por outra goroutine.
// Depois que o valor é lido do canal, ele é removido do canal, liberando o canal para receber um novo valor.
// Essa dinâmica de enviar e receber valores em um canal é o que permite a comunicação e a sincronização entre goroutines.

// Quando uma goroutine envia um valor para o canal, ela pode continuar sua execução sem esperar que outro processo leia o valor imediatamente.
// Da mesma forma, quando uma goroutine lê um valor do canal, ela bloqueia até que um valor esteja disponível para ser lido,
// garantindo que as operações sejam sincronizadas.
