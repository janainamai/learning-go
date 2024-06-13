package main

// demonstrando o uso do select em conjunto com o channel (chan)
import (
	"fmt"
	"time"
)

func main() {
	// criando um channel (canal) para transmitir uma string
	hello := make(chan string)

	// criando uma goroutine para esperar 2 segundos antes de enviar a mensagem "hello world" para o canal "hello"
	go func() {
		time.Sleep(time.Second * 2)
		hello <- "hello world"
	}()

	// criando um select para realizar operações de comunicação de forma não bloqueante
	// verifica se alguma das operações pode prosseguir imediatamente
	// se múltiplas operações estiverem prontas, uma delas é selecionada de forma automática
	select {
	// se houver um valor disponível no "hello", será atribuído para "x" e executando o println
	case x := <-hello:
		fmt.Println(x)
	// se não houver nenhum valor, o bloco default é executado
	default:
		fmt.Println("default")
	}

	// a saída será default pois a goroutine esperou 2 segundos para preencher o canal "hello"

	// mesmo que esperamos 5 segundos para finalizar a execução, o select já havia sido percorrido
	// por isso a saída continua sendo default
	// time.Sleep(time.Second * 5)
}
