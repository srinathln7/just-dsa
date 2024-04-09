package main

import "strconv"

func multiply(num1 string, num2 string) string {

	// Key Idea:  Reverse the two strings and start multiplying digit-by-digit
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	num1 = revStr(num1)
	num2 = revStr(num2)

	result := make([]int, m+n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			digit := int(num1[i]-'0') * int(num2[j]-'0')
			result[i+j] += digit
			result[i+j+1] += result[i+j] / 10
			result[i+j] %= 10
		}
	}
	revSlice(result)

	// Remove leading zeros
	beg := 0
	for beg < len(result) && result[beg] == 0 {
		beg++
	}

	// Convert the result back to string
	var res string
	for i := beg; i < len(result); i++ {
		res += strconv.Itoa(result[i])
	}

	return res
}

func revStr(s string) string {
	n := len(s)
	runes := []rune(s)
	for i := 0; i < n/2; i++ {
		// MISTAKE: s[i], s[n-1-i] = s[n-1-i], s[i] => String in Go is an IMMUTABLE slice of bytes.
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func revSlice(nums []int) []int {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
	return nums
}
