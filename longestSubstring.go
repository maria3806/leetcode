package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charCount := [128]int{}
	left := 0
	maxLength := 0

	for right, char := range s {
		charCount[char]++

		for charCount[char] > 1 {
			charCount[s[left]]--
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := "abcabcbb"
	result := lengthOfLongestSubstring(input)
	fmt.Println(result)
}
