/*
 46. Permutations https://leetcode.com/problems/permutations/
Medium
Given a collection of distinct integers, return all possible permutations.

Example:
Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
=>
									[1,2,3]
							/			|			\
						[1] [2,3]
					/		\
				[1,2] [3]	[1,3] [2]
				/				\
			[1,2,3] []		[1,3,2] []

The goal is to traverse down the tree with adding

*/
package main

import "fmt"

func main() {
	fmt.Println("Permutations for [1,2,3,4]: ", findPermutations([]int{1, 2, 3, 4}))
}

var R [][]int

func findPermutations(nums []int) [][]int {
	R = [][]int{}
	helper(nums, []int{})
	return R
}

func helper(nums []int, curr []int) {
	// fmt.Println(nums, curr, R)

	if len(nums) == 0 {
		cresult := []int{}
		cresult = append(cresult, curr...)
		R = append(R, cresult)
		return
	}

	for i := 0; i < len(nums); i++ {
		//add the current iteration of nums to curr
		curr = append(curr, nums[i])

		//we need to pass down the reminder of the nums[], for that we need to make a new copy of reminer of nums[]
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		//remove the currently added nums[i] from curr
		newNums = append(newNums[:i], newNums[i+1:]...)
		// newNums = append(newNums[:i], newNums[i+1:]...)
		helper(newNums, curr)
		curr = curr[:len(curr)-1]
	}
}
