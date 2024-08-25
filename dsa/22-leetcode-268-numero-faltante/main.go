package main

func missingNumber(nums []int) int {
	arrLength := len(nums)

	sum := arrLength * (arrLength + 1) / 2 // soma de um intervalo de números

	for _, value := range nums {
		sum -= value
	}

	return sum
}
