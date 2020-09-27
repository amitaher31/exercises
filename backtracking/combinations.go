/*
77. Combinations https://leetcode.com/problems/combinations/
Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
You may return the answer in any order.

Example 1:
Input: n = 4, k = 2
Output:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
=>
								n = 4, k=2
							n = [1, 2, 3, 4]
							/
					[1] [2,3,4]
				/		|
		[1,2] [3,4]	  [1,3] [2,4]
			/
	R =[[1,2]]
*/

package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
}

func combine(n int, k int) [][]int {
	helper(n, k, 1, []int{})
	return R
}

var R [][]int

func helper(n int, k int, idx int, curr []int) {
	if len(curr) == k {
		//add to result
		r1 := []int{}
		r1 = append(r1, curr...)
		R = append(R, r1)
		return
	}

	for i := idx; i <= n; i++ {
		curr = append(curr, i)
		helper(n, k, i+1, curr)
		curr = curr[:len(curr)-1]
	}
}
