package main

import "fmt"

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	previous := 0

	for i := 0; i < len(s); i++ {
		current := romanMap[s[i]]
		if current > previous {
			total += current - 2*previous
		} else {
			total += current
		}
		previous = current
	}
	return total
}

func main() {
	input := "MCMXCIV"
	result := romanToInt(input)
	fmt.Println(result)
}
