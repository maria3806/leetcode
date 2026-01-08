package main

import "fmt"

func twoSum(nums []int, target int) []int {
	indexByValue := map[int]int{}

	for currentIndex := 0; currentIndex < len(nums); currentIndex++ {
		currentValue := nums[currentIndex]
		neededValue := target - currentValue

		if foundIndex, exists := indexByValue[neededValue]; exists {
			return []int{foundIndex, currentIndex}
		}

		indexByValue[currentValue] = currentIndex
	}

	return []int{}
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(numbers, target)
	fmt.Println(result)
}
