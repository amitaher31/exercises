package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("For [1,2,3,4] largest time = ", largestTimeFromDigits([]int{1, 2, 3, 4}))
}

func largestTimeFromDigits(A []int) string {
	R = [][]int{}
	getPermutations(A, []int{})
	// fmt.Println("R = ", R)
	// fmt.Println("R = ", len(R))
	hrs, mins := -1, -1
	for _, v := range R {
		vHrs := v[0]*10 + v[1]
		vMins := v[2]*10 + v[3]

		if vHrs < 24 && vMins < 60 {
			if vHrs > hrs || (vHrs == hrs && vMins > mins) {
				hrs = vHrs
				mins = vMins
			}
		}
	}

	if hrs == -1 {
		return ""
	}

	sHrs, sMins := strconv.Itoa(hrs), strconv.Itoa(mins)
	if len(sHrs) < 2 {
		sHrs = "0" + sHrs
	}

	if len(sMins) < 2 {
		sMins = "0" + sMins
	}

	return sHrs + ":" + sMins
}

var R [][]int

func getPermutations(A []int, curr []int) {
	if len(curr) == 4 {
		r1 := []int{}
		r1 = append(r1, curr...)
		R = append(R, r1)
		return
	}

	for i := 0; i < len(A); i++ {
		curr = append(curr, A[i])
		newA := make([]int, len(A))
		copy(newA, A)
		newA = append(newA[:i], newA[i+1:]...)
		getPermutations(newA, curr)
		curr = curr[:len(curr)-1]
	}
}
