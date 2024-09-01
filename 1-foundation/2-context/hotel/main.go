package main

import (
	"context"
	"fmt"
	"time"
)

// aqui iremos simular um sistema de hotel onde o usuário realiza o pedido de uma reserva de quarto
// vamos simular o retorno de sucesso após 5 segundos, e vamos simular também o cancelamento do pedido por parte do usuário na "go func"

// temos um sistema de hotel ao qual possui um tempo limite de 5 segundos para dar uma resposta ao usuário,
// se em 5 segundos não obtivermos uma resposta, o usuário receberá uma mensagem de sucesso
// no caso de passar 5 segundos sem resposta, por trás dos panos o sistema inseriu o pedido de reserva em uma fila para ser processado

// nossa aplicação sempre terá um contexto principal em branco, chamado backgroud
// teremos 5 segundos para reservar o quarto, durante esse 5 segundos o usuário tem o tempo de cancelar a reserva, para simular isso usaremos o go func
func main() {

	// contexto principal da aplicação (background)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// no final do processo o contexto será finalizado (defer)
	defer cancel()

	// criamos essa thread para simular a ação do usuário de cancelar o pedido de reserva antes dos 5 segundos
	// se setarmos 10 segundos, não haverá cancelamento do pedido de reserva e obteremos a mensagem de pedido realizado com sucesso
	// se setarmos 4 segundos, o pedido será cancelado antes do sistema processar o pedido de reserva
	go func() {
		time.Sleep(time.Second * 4)
		cancel()
	}()

	reservarHotel(ctx)

}

// por convenção e boas práticas, contexto é sempre o primeiro atributo a ser passado na função
func reservarHotel(ctx context.Context) {

	// se eu ainda não sei o que fazer com o context, posso deixá-lo como TODO, dessa forma posso usar testar a aplicação mesmo sem ter a lógica do context ainda
	// context.TODO()

	// select fica guardando alguma condição, e quando a condição for verdadeira é executada a ação
	// pesquisar channels
	select {
	// quando o contexto for cancelado, deve imprimir "tempo excedido"
	case <-ctx.Done():
		fmt.Println("Tempo excedido para processar a requisição")
	// se passar de 5 segundos, significa que o pedido de reserva foi feito, apenas simulando o funcionamento do sistema
	case <-time.After(time.Second * 5):
		fmt.Println("Pedido realizado, em breve enviaremos no seu email a confirmação da reserva")
	}
}

/*
context.WithTimeout(ctx, time.Second*3) -> Configura um timeout no context
podemos configurar esse tipo de timeout em querys de database, e em requisições http

context.WithDeadline -> Cancela a request em x segundos caso ainda não tiver obtido uma resposta

context.WithValue(key, value) -> permite enviar uma key para posteriormente obter o conteúdo com context.Value
*/
