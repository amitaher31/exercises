package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
// 49. Group Anagrams
// Medium

// 3351

// 177

// Add to List

// Share
// Given an array of strings, group anagrams together.

// Example:

// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//   ["ate","eat","tea"],
//   ["nat","tan"],
//   ["bat"]
// ]
// Note:

// All inputs will be in lowercase.
// The order of your output does not matter.
// Submission - https://leetcode.com/submissions/detail/350952908/

func main() {
	i1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("Input 1 = ", i1)
	fmt.Println("result = ", groupAnagrams(i1))

	i2 := []string{"ray", "cod", "abe", "ned", "arc", "jar", "owl", "pop", "paw", "sky", "yup", "fed", "jul", "woo", "ado", "why", "ben", "mys", "den", "dem", "fat", "you", "eon", "sui", "oct", "asp", "ago", "lea", "sow", "hus", "fee", "yup", "eve", "red", "flo", "ids", "tic", "pup", "hag", "ito", "zoo"}
	fmt.Println("\nInput 2 = ", i2)
	fmt.Println("result = ", groupAnagrams(i2))
}

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)

	for _, v := range strs {
		sortedV := sortStr(v)
		mp[sortedV] = append(mp[sortedV], v)
	}

	r := [][]string{}
	for _, v := range mp {
		r = append(r, v)
	}

	return r
}

func sortStr(s string) string {
	sorted := []rune(s)
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i] < sorted[j] {
			return true
		}
		return false
	})
	return string(sorted)
}

// import "strings"
// func groupAnagrams(strs []string) [][]string {
//     r := [][]string{}

//     for _, s := range strs {
//         // fmt.Println("s = ", s)
//         if len(r) == 0 {
//             r = append(r, []string{s})
//             continue
//         }

//         add := false
//         for i, curr := range r {
//             add = false
//             // fmt.Println("curr = ", curr)
//             if len(s) != len(curr[0]) {
//                 add = true
//                 continue
//             }

//             for _, l := range s {
//                 // fmt.Println("l, curr = ", string(l), curr[0])
//                 if !strings.Contains(curr[0], string(l)) {
//                     // fmt.Println(" contains failed for l = ", string(l))
//                     add = true
//                     break
//                 }
//             }

//             if !add{
//                 curr1 := append(curr, s)
//                 r[i] = r[len(r)-1]
//                 r = r[:len(r)-1]
//                 r = append(r, curr1)
//                 add = false
//                 break
//             }
//         }

//         if add {
//             r = append(r, []string{s})
//         }
//     }

//     return r
// }
