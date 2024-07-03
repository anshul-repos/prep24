package main

import "fmt"

// Anagram strings are having same letters in different orders
// Ex1: "listen" and "silent"
// Ex2: "pot" and "otp"

func isAnagram(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	str1CharCountMap := make(map[rune]int)
	str2CharCountMap := make(map[rune]int)

	for _, char := range str1 {
		str1CharCountMap[char]++
	}

	for _, char := range str2 {
		str2CharCountMap[char]++
	}

	for char, count := range str1CharCountMap {
		if str2CharCountMap[char] != count {
			return false
		}
	}

	return true

}

func main() {
	str1 := "zaf"
	str2 := "qafz"
	fmt.Printf("%s and %s are Anagrams: %v", str1, str2, isAnagram(str1, str2))
	fmt.Println()
}
