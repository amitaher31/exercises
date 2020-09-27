package main

import "fmt"

func main() {
	fmt.Println(climbStairs_dp(12))
	fmt.Println(climbStairs(12))
	fmt.Println(climbStairs_recursive_memo(12))
}

// DP with bottom up approach => best & easiest to understand
func climbStairs_dp(n int) int {
	r := make([]int, n+1)
	r[0] = 1 //ans to smallest first problem
	r[1] = 1 //ans to smallest + 1 problem

	for i := 2; i <= n; i++ {
		r[i] = r[i-1] + r[i-2] //ans to the next problem = sum of last 2 problems
	}

	return r[n]
}

// Recursive only
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs_recursive_memo(n int) int {
	memo := make(map[int]int)
	return climb(n, memo)
}

func climb(n int, memo map[int]int) int {
	// fmt.Println("n = ", n)
	if n <= 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	v, ok := memo[n]
	if ok {
		// fmt.Println("ok := memo[n], v = ", v)
		return v
	}

	memo[n] = climbStairs(n-1) + climbStairs(n-2)
	return memo[n]
}
