/*
7. Reverse Integer https://leetcode.com/problems/reverse-integer/
Given a 32-bit signed integer, reverse digits of an integer.
Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Reverse of 123 = ", reverse(123))
}

func reverse(x int) int {
	// maxInt := math.MaxInt32
	negative, overflow := false, false
	rev, curr := 0, 0

	if x < 0 {
		negative = true
		x = -x
	}

	for x > 0 {
		curr = x % 10
		rev = rev*10 + curr
		if rev > math.MaxInt32 {
			overflow = true
			break
		}

		x = x / 10
	}

	if overflow {
		return 0
	}

	if negative {
		rev = -rev
	}

	return rev
}
