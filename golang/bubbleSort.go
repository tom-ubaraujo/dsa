package main

import "fmt"

func bSort(nums []int) []int {
	swap := true
	temp := 0

	for swap {

		swap = false
		for i := 1; i <= len(nums)-1; i++ {

			if nums[i-1] > nums[i] {
				temp = nums[i-1]
				nums[i-1] = nums[i]
				nums[i] = temp
				swap = true
			}
		}
	}
	return nums
}

func main() {
	test := []int{9, 4, 1, 6, 10, 2, 3, 7, 8, 5}
	fmt.Println(bSort(test))
}

/*

O LOOP WHILE FICA ATË QUE NÃO HAJA NENHUMA TROCA NO FOR
DENTRO DO WHILE VAMOS FAZER FOR NO ARRAY E COMPARAR OS NUMEROS DE 2 EM 2
SE TROCAR, SETAMOS SWAP PARA TRUE, SE NÃO APENAS SEGUE O LOOP

APENAS SAIMOS DO WHILE QUANDO FIZERMOS UM LOOP COMPLETO NO ARRAY SEM SWAP

*/
