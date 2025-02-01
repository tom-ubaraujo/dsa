// CONTAINS DUPLICATE

package main

import "fmt"

func main() {
	arr1 := []int{1, 5, 7, 10, 4, 1}
	arr2 := []int{2, 5, 7, 10, 4, 1}
	arr3 := []int{3, 7, 7, 10, 0, 9}

	fmt.Println(hasDupOn(arr1))
	fmt.Println(hasDupOn(arr2))
	fmt.Println(hasDupOn(arr3))

	fmt.Println(hasDupN2(arr1))
	fmt.Println(hasDupN2(arr2))
	fmt.Println(hasDupN2(arr3))
}

// o(n) solution
func hasDupOn(arr []int) bool {
	visited := make(map[int]bool)

	for _, v := range arr {
		if _, ok := visited[v]; ok {
			return true
		}
		visited[v] = true
	}
	return false
}

// o(n^2) solution
func hasDupN2(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for i2 := i + 1; i2 < len(arr); i2++ {
			if arr[i] == arr[i2] {
				return true
			}
		}
	}
	return false
}
