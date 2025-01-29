package main

import "fmt"

func main() {
	arr := []int{0, 2, 10, 0, 0, 3, 0, 5, 6, 2, 0, 0, 0, 0, 1, 4, 2, 0, 0}
	result := moveZerosToEnd(arr)
	fmt.Println(result)
}

func moveZerosToEnd(arr []int) []int {
	j := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		if arr[j] != 0 {
			j++
		}
	}

	return arr
}
