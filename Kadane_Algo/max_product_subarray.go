/*
152. Maximum Product Subarray https://leetcode.com/problems/maximum-product-subarray/
Medium
Given an integer array nums, find the contiguous subarray within an array (containing at least one number) which has the largest product.

Example 1:

Input: [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.
Example 2:

Input: [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
*/
package main

import "fmt"

func main() {
	Arr := [][]int{}
	Arr = append(Arr, []int{1, 2, 3, 4})
	Arr = append(Arr, []int{-1, -2, -9, -6})
	Arr = append(Arr, []int{2, 3, -2, 4})
	Arr = append(Arr, []int{-2, 0, -1})
	Arr = append(Arr, []int{-2})

	for _, a := range Arr {
		fmt.Println("Max prod sub-array for ", a, maxProduct(a))
	}
}
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar, minSoFar := nums[0], nums[0]
	result := maxSoFar
	// fmt.Println(maxSoFar, minSoFar, result)

	for i, curr := range nums {
		if i == 0 {
			// since use the nums[0] value as init
			continue
		}
		// cases: from which we need to find max
		// maxSoFar was 0, so curr could be increase or decrease maxSoFar if curr is +ve/-ve
		// maxSoFar is +ve, so maxSoFar * curr would increase maxSoFar if curr is +ve
		// maxSoFar is -ve, so maxSoFar * curr would decrease or increase based on if curr is +ve/-ve
		tempMax := max(curr, max(maxSoFar*curr, minSoFar*curr))

		minSoFar = min(curr, min(maxSoFar*curr, minSoFar*curr))

		maxSoFar = tempMax

		result = max(maxSoFar, result)

		// fmt.Println(maxSoFar, minSoFar, result)
	}

	return result
}

// [-2,0,-1]
// curr = -2 maxSofar = 0, minsofar = -2, r = 0
// curr = 0, maxsofar = 0, minsofar = 0, r = 0
// curr = -1 maxsofar = 0, 0

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// brute force
// func maxProduct(nums []int) int {
//     max := 0

//     for i:=0; i<len(nums); i++ {
//         iproduct := nums[i]
//         for j:=i+1; j<len(nums); j++ {
//             iproduct = iproduct * nums[j]
//             if max < iproduct {
//                 max = iproduct
//             }
//         }
//     }

//     return max
// }
