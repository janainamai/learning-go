package main

import (
	"fmt"
)

// cooperative scheduling
// aqui tentaremos visualizar como funciona a forma cooperativa das threads trabalharem no Go

// este modelo de concorrência veio na versão 1.14,
// refere-se à abordagem em que as goroutines cooperam voluntariamente para compartilhar o tempo de CPU e permitir que outras tarefas sejam executadas
// isso é feito sem a interferência direta do sistema operacional ou de um mecanismo de escalonamento externo
// em vez disso, as goroutines cooperam entre si, liberando voluntariamente o controle da CPU quando necessário
// e evitando bloqueios desnecessários.

// nesse exemplo nosso programa deverá imprimir "começou", entrar em looping infinito e nunca mais encerrar o programa
// porém com a funcionalidade de cooperação entre as threads, o schedule entende que tem uma thread sendo usada de forma desnecessária
// com isso ele libera a thread para continuar a execução, imprimindo "terminou" e finalizando o programa
func main() {

	fmt.Println("Começou")

	go func() {
		for {

		}
	}()

	fmt.Println("Terminou")
}
