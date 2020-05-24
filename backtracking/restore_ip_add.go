/*
https://leetcode.com/problems/restore-ip-addresses/
Given a string containing only digits, restore it by returning all possible valid IP address combinations.

A valid IP address consists of exactly four integers (each integer is between 0 and 255) separated by single points.

Example:

Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
*/
package exercises

import (
	"fmt"
	"strconv"
	"strings"
)

var result = []string{}

func restoreIpAddresses(s string) []string {
	r := []string{}
	result = []string{}

	helper(s, r, 0)
	return result
}

//25525511135
// 2.5.5.255555*
func helper(s string, r []string, start int) {
	fmt.Printf("s = %+v, r = %+v \n", s, r)
	if len(r) == 4 && start == len(s) {
		result = append(result, strings.Join(r, "."))
		return
	}

	if len(r) >= 4 {
		return
	}

	for i := start; i < len(s); i++ {
		s0 := s[start : i+1]
		if isValid(s0) {
			helper(s, append(r, s0), i+1)
		} else {
			break
		}
	}
}

func isValid(s0 string) bool {

	if s0[0] == '0' && len(s0) > 1 {
		return false
	}

	s0Int, err := strconv.Atoi(s0)
	if err != nil {
		return false
	}

	if s0Int == 0 && len(s0) > 1 {
		return false
	}

	if s0Int <= 255 && s0Int >= 0 {
		return true
	}
	return false
}

// 25525511135
// 2 5525511135
//   5 5 25511135
