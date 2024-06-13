package main

import "fmt"

func main() {

	// variáveis são alocadas em endereços de memórias
	// ponteiros apontam para endereços de memória
	// ponteiros são variáveis que alocam o endereço de memória de uma variável
	// & retorna o endereço de uma variável
	// * usado para declarar ponteiros ou para acessar o valor armazenado no endereço de memória apontado por um ponteiro

	// essa variável, quando criada, é alocada em um lugar na memória do computador
	// em outras palavras, a variável "numero" possui um endereço em memória
	numero := 10

	// obtendo o endereço em memória da variável "a"
	fmt.Println(&numero)

	// criando um ponteiro para o endereço de memória da variável "numero"
	// agora temos um ponteiro, essa variável possui como valor o endereço em memória da variável "numero", ex: 0xc000094018
	var ponteiro *int = &numero

	// obtendo o valor contido no endereço em memória da variável "numero"
	valorNumero := *ponteiro
	fmt.Println(valorNumero)

	fmt.Print("Endereço em memória da variável ponteiro:", ponteiro) // será igual ao endereço da variável "numero"
	fmt.Printf("\nValor da variável ponteiro: %d \n", *ponteiro)     // 10

	// se eu quiser alterar o valor que está alocado em um endereço de memória, basta inserir o asterisco na frente da variável
	*ponteiro = 50
	fmt.Println(numero) // 50

	// testando mais exemplos

	// declarando variável "nome" com valor "Janaina"
	nome := "Janaina"

	// declarando variável "aluna" atribuindo o endereço de memória da variável "nome"
	aluna := &nome

	fmt.Printf("\n\nNome: %v \nAluna: %v \n\n", nome, *aluna)
	fmt.Print(&nome, aluna)

	// alterando o valor do endereço de memória de ambas as variáveis, pois é o mesmo
	*aluna = "Janaina Mai"

	fmt.Printf("\n\nNome: %v \nAluna: %v \n\n", nome, *aluna)
	fmt.Print(&nome, aluna)

	// podemos utilizar ponteiros em funções
	// o método mudarNome espera receber um ponteiro, então para enviar "nome" preciso utilizar & e para enviar aluna não, pois aluna é um ponteiro
	mudarNome(&nome)
	mudarNome(aluna)

	fmt.Println("\n", nome)

}

func mudarNome(nome *string) {
	*nome = "Fulana"
}
