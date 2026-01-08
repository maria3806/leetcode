package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := &strings.Builder{}

	for index, value := range values {
		for num >= value {
			num -= value
			result.WriteString(symbols[index])
		}
	}
	return result.String()
}

func main() {
	number := 1994
	result := intToRoman(number)
	fmt.Println(result)
}
