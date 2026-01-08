package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	result := findMedianSortedArrays(nums1, nums2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenA := len(nums1)
	lenB := len(nums2)

	var findKth func(indexA, indexB, kth int) int

	findKth = func(indexA, indexB, kth int) int {
		if indexA >= lenA {
			return nums2[indexB+kth-1]
		}
		if indexB >= lenB {
			return nums1[indexA+kth-1]
		}
		if kth == 1 {
			return min(nums1[indexA], nums2[indexB])
		}

		half := kth / 2
		maxInt := int(^uint(0) >> 1)

		valueA := maxInt
		valueB := maxInt

		if indexA+half-1 < lenA {
			valueA = nums1[indexA+half-1]
		}
		if indexB+half-1 < lenB {
			valueB = nums2[indexB+half-1]
		}

		if valueA < valueB {
			return findKth(indexA+half, indexB, kth-half)
		}
		return findKth(indexA, indexB+half, kth-half)
	}

	leftMedian := findKth(0, 0, (lenA+lenB+1)/2)
	rightMedian := findKth(0, 0, (lenA+lenB+2)/2)

	return float64(leftMedian+rightMedian) / 2.0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
