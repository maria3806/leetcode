package main

import (
	"fmt"
	"math"
)

func reverse(number int) int {
	reversed := 0

	for number != 0 {
		if reversed < math.MinInt32/10 || reversed > math.MaxInt32/10 {
			return 0
		}
		reversed = reversed*10 + number%10
		number /= 10
	}
	return reversed
}

func main() {
	number := 123
	result := reverse(number)
	fmt.Println(result)
}
