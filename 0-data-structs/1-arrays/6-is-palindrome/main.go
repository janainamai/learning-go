package main

import "fmt"

func main() {
	printIsPalindrome("ARARA")
	printIsPalindrome("HELLO")
}

func printIsPalindrome(word string) {
	fmt.Printf("The word '%s' is palindrome: %v\n", word, isPalindrome(word))
}

func isPalindrome(word string) bool {
	firstLetter := 0
	lastLetter := len(word) - 1

	for firstLetter < lastLetter {

		if word[firstLetter] != word[lastLetter] {
			return false
		}

		firstLetter++
		lastLetter--
	}

	return true
}
