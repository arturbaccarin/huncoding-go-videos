package main

func moveZeroes(arr []int) []int {
	j := 0
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		if arr[j] != 0 {
			j++
		}
	}

	return arr
}

func moveZeros(arr []int) []int {
	arrLen := len(arr)

	for i := 0; i < arrLen; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i], arr[i+1:]...)
			arr = append(arr, 0)
			i -= -1
			arrLen -= 1
		}
	}

	return arr
}
