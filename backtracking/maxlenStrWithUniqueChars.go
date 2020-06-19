import "fmt"

/*
1239. Maximum Length of a Concatenated String with Unique Characters
https://leetcode.com/problems/maximum-length-of-a-concatenated-string-with-unique-characters/
Given an array of strings arr. String s is a concatenation of a sub-sequence of arr which have unique characters.

Return the maximum possible length of s
Example 1:

Input: arr = ["un","iq","ue"]
Output: 4
Explanation: All possible concatenations are "","un","iq","ue","uniq" and "ique".
Maximum length is 4.
Example 2:

Input: arr = ["cha","r","act","ers"]
Output: 6
Explanation: Possible solutions are "chaers" and "acters".
Example 3:

Input: arr = ["abcdefghijklmnopqrstuvwxyz"]
Output: 26


Constraints:

1 <= arr.length <= 16
1 <= arr[i].length <= 26
arr[i] contains only lower case English letters.
Accepted
22,950
Submissions
48,514
*/
func maxLength(arr []string) int {
	if len(arr) == 0 {
		return 0
	}

	Arr = arr
	r = 0
	backtrack("", 0)
	return r
}

var Arr []string
var r int

func backtrack(curr string, idx int) {
	fmt.Println(curr, idx)
	if isUniqueChars(curr) {
		r = max(r, len(curr))
	} else {
		return
	}

	if idx == len(Arr) {
		return
	}

	for i := idx; i < len(Arr); i++ {
		backtrack(curr+Arr[i], i+1)
	}

}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func isUniqueChars(s string) bool {
	alph := make([]bool, 26)
	for _, v := range s {
		ch := v - 'a'
		if alph[ch] == true {
			return false
		}
		alph[ch] = true
	}
	return true
}