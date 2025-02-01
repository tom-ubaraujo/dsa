// is anagram

package main

import "fmt"

func main() {

	str1 := "asdfg"
	str2 := "gfdsa"

	str3 := "asdfg"
	str4 := "ajaui"

	fmt.Println(isAnagram(str1, str2))
	fmt.Println(isAnagram(str3, str4))
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charCounts := make([]int, 26)

	for i := 0; i < len(s); i++ {
		charCounts[s[i]-'a']++
		charCounts[t[i]-'a']--
	}

	fmt.Println(charCounts)

	for _, v := range charCounts {
		if v != 0 {
			return false
		}
	}
	return true
}
