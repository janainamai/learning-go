package main

import "fmt"

// Um channel (canal) em Go é uma estrutura de dados que proporciona uma forma de comunicação e sincronização entre goroutines, permitindo que elas troquem dados de forma segura e coordenada.

// Para entender a utilidade do channel, imagine 2 threads trabalhando e modificando o mesmo endereço de memória da variável A.
// `thread 1-sqlpuro seta A= 10 e depois envia um email`
// `thread 2 seta A= 20 e depois envia um email`
// Se a thread 2 setar o valor de A igual a 20 antes da threar 1-sqlpuro enviar o email, o fluxo ficará inconsistente

// Então o channel evita que as threads acessem memórias em comum para se comunicar, ele faz com que as threads se comuniquem entre si para evitar a concorrência na memória.

// Para visualizar esse exemplo, precisamos entender que main é uma thread, e a goroutine é outra thread
// THREAD 1-sqlpuro
func main() {

	// criando um canal para a string "hello"
	hello := make(chan string)

	// criando goroutine que executa uma função anônima para alterar o valor de hello
	// THREAD 2
	go func() {
		hello <- "Hello world"
	}()

	// configurando para que, sempre que a variável hello tiver um valor, enviar para a variável result
	// essa parte do código só será executada quando a gorountine atribuir um valor para a variável "hello"
	result := <-hello
	fmt.Println(result)

	// se eu quiser que o sistema aguarde para finalizar até que a variável hello receba um valor:
	// <-hello
}
