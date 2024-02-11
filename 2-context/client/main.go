package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// aqui vamos fazer uma requisição http enviando um contexto configurado para retornar timeout após 6 segundos sem resposta
// simulando o cancelamento de uma request por parte do client http

// executar main.go da pasta server, que possui uma rota http:8080 que retorna sucesso em 5 segundos
// executar main.go da pasta client, que faz requisição GET para localhost:8080 e espera 6 segundos para dar timeout

// se configurarmos o timeout para 3 segundos:
// o que acontecerá no client: não teremos a resposta de sucesso, pois em 3 segundos o contexto retornará timeout e será encerrado, mostrando o erro "context deadline exceeded" no terminal
// o que acontecerá no server: no terminal do server será possível visualizar que após a request ser cancelada pelo client, o servidor interromperá o processamento da request
func main() {
	// configurando um contexto para retornar timeout após 6 segundos sem resposta
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// preparando request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// executando chamada http
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// fechando a requisição
	defer res.Body.Close()

	// setando o conteúdo do body do resultado da request para aparecer no terminal
	io.Copy(os.Stdout, res.Body)

}
