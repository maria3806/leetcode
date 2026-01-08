package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	index := 0
	length := len(s)
	result := 0

	for index < length && s[index] == ' ' {
		index++
	}

	if index == length {
		return 0
	}

	signMultiplier := 1
	if s[index] == '-' {
		signMultiplier = -1
		index++
	} else if s[index] == '+' {
		index++
	}

	for index < length && s[index] >= '0' && s[index] <= '9' {
		result = result*10 + int(s[index]-'0')
		index++

		if result > math.MaxInt32 {
			break
		}
	}

	if result > math.MaxInt32 {
		if signMultiplier == -1 {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	return signMultiplier * result
}

func main() {
	input := "-42"
	result := myAtoi(input)
	fmt.Println(result)
}
