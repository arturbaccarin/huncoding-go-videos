/*
Array

Tamanho fixo e não pode ser alterado
Capa partição é indexada.
Coleção de dados de tum tipo específico

Todas as posições tem sua memória contígua.

Bloco de memória consecutivos
*/
package main

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
