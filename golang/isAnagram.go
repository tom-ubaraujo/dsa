package main

import (
	"fmt"
)

func anagram(a, b string) bool {
	if a == b {
		return true
	}

	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)

	for _, i := range a {
		m[i] += 1
	}

	for _, i := range b {
		_, ok := m[i]
		if ok {
			m[i] -= 1
			if m[i] == 0 {
				delete(m, i)
			}
		} else {
			return false
		}
	}

	if len(m) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(anagram("banana", "nabana"))
	fmt.Println(anagram("bacana", "cabana"))
	fmt.Println(anagram("ana", "ana"))
	fmt.Println(anagram("teste", "tomzera"))
	fmt.Println(anagram("meat", "team"))
	fmt.Println(anagram("conservadora", "conversadora"))
	fmt.Println(anagram("alegria", "alergia"))
	fmt.Println(anagram("players", "parsley"))
	fmt.Println(anagram("eleven plus two", "twelve plus one"))
}
