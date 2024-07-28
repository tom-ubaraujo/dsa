package main

func linear_search(myArr []int, v int) bool {
	for i := range myArr {
		if i == v {
			return true
		}
	}
	return false
}

func main() {
	arrTest := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrTest2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(linear_search(arrTest, 5))   // return true
	println(linear_search(arrTest2, 11)) // return false
}
