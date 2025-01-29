package main

import "fmt"

// in this example, we have 9 numbers, zero is hidding a number between 1 and 9
// in this case, the hidden number is 3
func main() {
	arr := []int{4, 5, 6, 1, 2, 8, 9, 0, 7}

	hiddenNumber := getHiddenNumber(arr)

	fmt.Println(hiddenNumber)
}

func getHiddenNumber(arr []int) int {
	arrLenght := len(arr)
	sum := arrLenght * (arrLenght + 1) / 2

	for _, value := range arr {
		sum -= value
	}

	return sum
}
