package main

import (
	"log"
	"net/http"
	"time"
)

// simulando uma requisição cancelada do lado do servidor, ou seja, após enviar a requisição, cancelar ela antes de obter resposta
// executar arquivo main, isso irá iniciar o servidor
// chamar rota localhost:8080
// dar ctrl + c antes de obter resposta, a requisição deve ser cancelada
// aguardar resposta, deve retornar sucesso
func main() {

	// configurando rota "/" para chamar função "home"
	http.HandleFunc("/", home)

	// iniciando server na porta 8080
	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou a request")
	defer log.Println("Finalizou minha request")

	select {
	// simulando que a requisição é processada com sucesso em 5 segundos
	case <-time.After(time.Second * 5):
		log.Println("Requisição processada com sucesso")
		w.Write([]byte("Requisição processada com sucesso")) // escreve a mensagem de retorno do endpoint

	// se o client cancelar a requisição em menos de 5 segundos, devo parar a requisição
	case <-ctx.Done(): // acontece se eu der um ctrl + c para cancelar a request após chamar a rota localhost:8080/
		log.Println("Request cancelada")
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)
	}
}
