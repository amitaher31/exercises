/*
763. Partition Labels (Medium) https://leetcode.com/problems/partition-labels/
A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.
Example 1:
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.

Note:
S will have length in range [1, 500].
S will consist of lowercase English letters ('a' to 'z') only.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("aay"))
}

/*
Solution: The goal is to partition the str into as many parts possible that each letter appears in 1 part
- so we need to make sure all letters in a part is in a partition
- if we have last occurance for each letter in S
- we can try to create a partition of S such that S[i to j] would span from 1 to last occurance of i to j*/

func partitionLabels(S string) []int {
	R := []int{}
	if len(S) == 0 {
		return R
	}

	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[c-'a'] = i
	}

	i := 0
	for i < len(S) {
		end := lastIndex[S[i]-'a']
		j := i
		for j < end {
			end = max(end, lastIndex[S[j]-'a'])
			j++
		}
		R = append(R, j-i+1)
		i = j + 1
	}

	return R
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
