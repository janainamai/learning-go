package main

// Função que usa um loop para calcular a soma
func SumLoop(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// Função que usa recursão para calcular a soma
func SumRecursion(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + SumRecursion(nums[1:])
}
