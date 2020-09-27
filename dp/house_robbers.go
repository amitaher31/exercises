/*
## 198. House Robber,
https://leetcode.com/problems/house-robber/
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
Example 1:
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.

## 213. House Robber II
https://leetcode.com/problems/house-robber-ii/
 All houses at this place are arranged in a circle.
 That means the first house is the neighbor of the last one

 Example 1:
Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
			 because they are adjacent houses.

Example 2:
Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
*/

/*
=>>
The goal is to make sure pick the houses which maximises our stash
so if we have 1 house, we can pick 1
if we have 2 houses, max(1,2)
if 3 - max(1+3, 2)
if 4 - max(max(1,2)+4, max(1+3, 2))
if k, max(f(k-2)+nums[k], f(k-1))
*/

package main

import "fmt"

func main() {
	fmt.Println("House Robber I")
	fmt.Println("DP input = [1, 2, 3, 1], result = ", rob_dp([]int{1, 2, 3, 1}))
	fmt.Println("DP input = [1, 2, 3, 1], result = ", rob_dp([]int{2, 7, 9, 3, 1}))

	fmt.Println("DP input = [1, 2, 3, 1], result = ", rob([]int{1, 2, 3, 1}))
	fmt.Println("DP input = [1, 2, 3, 1], result = ", rob([]int{2, 7, 9, 3, 1}))
	fmt.Println("-----------------\n")

	fmt.Println("House Robber II")
	fmt.Println("input = [2,3,2], result = ", rob_circle([]int{2, 3, 2}))
	fmt.Println("input = [1,2,3,1], result = ", rob_circle([]int{1, 2, 3, 1}))
	fmt.Println("-----------------\n")
}

//dp -> uses extra space for dp array
func rob_dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}

func rob(nums []int) int {
	last_max, prev_last_max := 0, 0
	for i := 0; i < len(nums); i++ {
		t := last_max
		last_max = max(nums[i]+prev_last_max, last_max)
		prev_last_max = t
	}
	return last_max
}

/*
When the houses are arranged in a circle, you cannot rob the 1st & last house together (since they can be termed as adjecent)
- the idea then is to pick either of them by dividing the circle in 2 partitions, so the problem becomes 2 sets of hour robber 1 problem & we can take the max of them
  partition 1 = house 0 to 2nd last house (exclude last)
  partition 2 = house 1 to last house (exclude first)
  return max(rob(part1), rob(part2))
*/
func rob_circle(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[1]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	//0 to 2nd to last
	part1 := nums[:len(nums)-1]
	// fmt.Println("part1 = ", part1)

	//1 to last
	part2 := nums[1:]
	// fmt.Println("part2 = ", part2)

	return max(rob(part1), rob(part2))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rob3(root *TreeNode) int {
	s1, s2 := helper(root)
	return max(s1, s2)
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	sum := 0

	sum += root.Val
	lsum1, lsum2 := helper(root.Left)
	rsum1, rsum2 := helper(root.Right)

	max1 := max(lsum1, lsum2)
	max2 := max(rsum1, rsum2)

	return sum + lsum2 + rsum2, max1 + max2
}
