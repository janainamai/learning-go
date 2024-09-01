package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// a maioria das funções em Go seguem esse padrão de retornar algo e também um error
	resp, err := http.Get("http://google.com.br")

	// é comum verificar o err antes de prosseguir
	if err != nil {
		// tratar o erro da melhor forma para o seu sistema, conforme regras de negócio, etc
		log.Fatal(err.Error())
	}

	fmt.Println(resp.StatusCode)

	// testando o retorno da função criada, deve retornar erro
	result, err := soma(5, 9)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(result)

	// se eu não quiser utilizar o erro, posso ignorar esse retorno da função com _
}

// como retornar um erro em uma função
func soma(x int, y int) (int, error) {
	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}
