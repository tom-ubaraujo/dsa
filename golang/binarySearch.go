package main

func binaryearch(myArr []int, v int) bool {
	low, hi := 0, len(myArr)

	for low < hi {
		mid := (hi + low) / 2

		if myArr[mid] == v {
			return true
		} else if myArr[mid] > v {
			hi = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	arrTest := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrTest2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	println(binaryearch(arrTest, 5))   // return true
	println(binaryearch(arrTest2, 11)) // return false
}
