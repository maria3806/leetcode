package main

import "fmt"

func isMatch(s string, p string) bool {
	lenS := len(s)
	lenP := len(p)

	memo := make([][]int, lenS+1)
	for i := range memo {
		memo[i] = make([]int, lenP+1)
	}

	var matchHelper func(indexS, indexP int) bool
	matchHelper = func(indexS, indexP int) bool {
		if indexP >= lenP {
			return indexS == lenS
		}

		if memo[indexS][indexP] != 0 {
			return memo[indexS][indexP] == 1
		}

		result := -1

		if indexP+1 < lenP && p[indexP+1] == '*' {
			if matchHelper(indexS, indexP+2) ||
				(indexS < lenS && (s[indexS] == p[indexP] || p[indexP] == '.') && matchHelper(indexS+1, indexP)) {
				result = 1
			}
		} else if indexS < lenS && (s[indexS] == p[indexP] || p[indexP] == '.') && matchHelper(indexS+1, indexP+1) {
			result = 1
		}

		memo[indexS][indexP] = result
		return result == 1
	}
	return matchHelper(0, 0)
}

func main() {
	s := "aab"
	p := "c*a*b"

	result := isMatch(s, p)
	fmt.Println(result)
}
