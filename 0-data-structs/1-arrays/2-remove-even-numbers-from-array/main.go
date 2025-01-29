package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	result := removeEvenNumbersFromArray(arr)

	fmt.Println(result)
}

func removeEvenNumbersFromArray(arr []int) []int {
	oddNumbers := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddNumbers = append(oddNumbers, arr[i])
		}
	}

	return oddNumbers
}
