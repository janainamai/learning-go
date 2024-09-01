package main

import (
	"fmt"
	"time"
)

// criando uma função para ler o canal msg
func worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker: ", workerId, " mensagem: ", res)
		time.Sleep(time.Second)
	}
}

// tendo em vista que o worker vai iterar o canal até chegar no número 10
// criamos 2 workers que são na verdade 2 threads pois criamos 1-sqlpuro goroutine para cada worker
// essas 2 threads chamaram o método "worker" para iterar até 10
// quanto mais workers criarmos, mais rápido o for chega até o 10
func main() {
	msg := make(chan int)
	go worker(1, msg) // worker 1-sqlpuro
	go worker(2, msg) // worker 2

	for i := 0; i < 10; i++ {
		msg <- i
	}

	// resultado usando somente worker 1-sqlpuro:
	// Worker:  1-sqlpuro  mensagem:  0
	// Worker:  1-sqlpuro  mensagem:  1-sqlpuro
	// Worker:  1-sqlpuro  mensagem:  2
	// Worker:  1-sqlpuro  mensagem:  3
	// Worker:  1-sqlpuro  mensagem:  4
	// Worker:  1-sqlpuro  mensagem:  5
	// Worker:  1-sqlpuro  mensagem:  6

	// resultado usando somente worker 2:
	// Worker:  2  mensagem:  1-sqlpuro
	// Worker:  1-sqlpuro  mensagem:  0
	// Worker:  1-sqlpuro  mensagem:  2
	// Worker:  2  mensagem:  3
	// Worker:  2  mensagem:  4
	// Worker:  1-sqlpuro  mensagem:  5
	// Worker:  1-sqlpuro  mensagem:  6
	// Worker:  2  mensagem:  7
	// Worker:  1-sqlpuro  mensagem:  8
	// Worker:  2  mensagem:  9
}
