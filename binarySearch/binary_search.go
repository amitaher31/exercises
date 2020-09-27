package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("nums = ", nums)
	fmt.Println("find index for 2: ", search(nums, 2))
	fmt.Println("find index for 6: ", search(nums, 6))
	fmt.Println("find index for 7: ", search(nums, 7))
	fmt.Println("\nums = ", nums)
	nums = []int{0, 1}
	fmt.Println("find index for 0: ", search(nums, 0))
	fmt.Println("find index for 1: ", search(nums, 1))
	nums = []int{0}
	fmt.Println("\nnums = ", nums)
	fmt.Println("find index for 0: ", search(nums, 0))
	fmt.Println("find index for 1: ", search(nums, 1))
	// fmt.Println(mySqrt(8))
	fmt.Println("\nFind Square Root")
	fmt.Println("6 = ", mySqrt(36))
	fmt.Println("38 = ", mySqrt(38))
	fmt.Println("8 = ", mySqrt(8))

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("\nFind in Rotated Array")
	fmt.Println("0 = ", searchRotated(nums, 0))
	fmt.Println("3 = ", searchRotated(nums, 3))
}

func search(nums []int, target int) int {

	a, b := 0, len(nums)-1
	for a <= b {
		mid := a + (b-a)/2
		// fmt.Println(a, b, mid)

		if nums[mid] == target {
			// fmt.Println("foudn")
			return mid
		} else if target < nums[mid] {
			// fmt.Println("setting b")
			b = mid - 1
		} else {
			// fmt.Println("setting a")
			a = mid + 1
		}

	}
	return -1
}

func mySqrt(x int) int {

	if x < 2 {
		return x
	}
	a, b := 2, x/2

	for a <= b {
		m := a + (b-a)/2
		// fmt.Println(a,b,m)
		mm := m * m
		if mm > x {
			b = m - 1
		} else if mm < x {
			a = m + 1
		} else {
			return m
		}
	}

	return b

}

func searchRotated(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	a, b := 0, len(nums)-1
	for a <= b {
		mid := a + (b-a)/2
		if target == nums[mid] {
			return mid
		}

		if nums[mid] >= nums[a] {
			if target >= nums[a] && target < nums[mid] {
				b = mid - 1
			} else {
				a = mid + 1
			}
		} else {
			if target <= nums[b] && target > nums[mid] {
				a = mid + 1
			} else {
				b = mid - 1
			}

		}
	}
	return -1
}
