package main

import "fmt"

func main() {
	arr := []int{7, 18, 5, 23, 1}

	result := reverseArray(arr)

	fmt.Println(result)
}

func reverseArray(arr []int) []int {
	start := 0
	end := len(arr) - 1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
