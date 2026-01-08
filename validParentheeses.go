package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		last := stack[len(stack)-1]

		if !isPair(last, char) {
			return false
		}

		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func isPair(open, close rune) bool {
	if open == '(' && close == ')' {
		return true
	}
	if open == '[' && close == ']' {
		return true
	}
	if open == '{' && close == '}' {
		return true
	}
	return false
}

func main() {
	input := "()[]{}"
	result := isValid(input)
	fmt.Println(result)
}
