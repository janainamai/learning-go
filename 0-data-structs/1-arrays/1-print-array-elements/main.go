package main

import "fmt"

func main() {
	var arr = []int{21, 3, 432, 5543, 5643}

	printArray(arr)
}

func printArray(arr []int) {
	for _, value := range arr {
		fmt.Printf("%d\n", value)
	}
}
