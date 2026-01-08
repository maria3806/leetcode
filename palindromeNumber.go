package main

import "fmt"

func isPalindrome(number int) bool {
	if number < 0 || (number > 0 && number%10 == 0) {
		return false
	}

	reversedHalf := 0

	for reversedHalf < number {
		reversedHalf = reversedHalf*10 + number%10
		number /= 10
	}
	return number == reversedHalf || number == reversedHalf/10
}

func main() {
	input := 121
	result := isPalindrome(input)
	fmt.Println(result)
}
