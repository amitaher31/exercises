/*
https://leetcode.com/problems/letter-combinations-of-a-phone-number/
Leet code 17
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

Analysis
map of digit to [a b c]
- from digits form a list of digit letters eg: 23 => dl =  [[abc], [def]]
- result = []strings
- ir = []string => intermediate result
start with 2 digits
- call helper to create permutation for each ele in dl [] & store in ir
- helper(dl, c, t) where c, t => current & target dl elements to run permutation against
  so dl runs for i:=0 to i < len(dl[t]){
      ir =
  }
*/

package main

import "fmt"

func main() {
	fmt.Println("23 = ", letterCombinations("23"))
	fmt.Println("2 = ", letterCombinations("2"))
	fmt.Println("nil = ", letterCombinations(""))
	fmt.Println("237 = ", letterCombinations("237"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	phmap := map[string][]string{}
	phmap["2"] = []string{"a", "b", "c"}
	phmap["3"] = []string{"d", "e", "f"}
	phmap["4"] = []string{"g", "h", "i"}
	phmap["5"] = []string{"j", "k", "l"}
	phmap["6"] = []string{"m", "n", "o"}
	phmap["7"] = []string{"p", "q", "r", "s"}
	phmap["8"] = []string{"t", "u", "v"}
	phmap["9"] = []string{"w", "x", "y", "z"}

	dl := [][]string{}
	for _, v := range digits {
		dl = append(dl, phmap[string(v)])
	}

	if len(dl) < 2 {
		return dl[0]
	}

	// fmt.Println("dl = ", dl)
	helper(dl, 0, 1)
	return dl[len(dl)-1]
}

func helper(dl [][]string, c int, t int) {
	// fmt.Println("dl = ", dl, c, t)
	if len(dl) == t {
		return
	}

	l1, l2 := dl[c], dl[t]
	tempR := []string{}
	for _, l2v := range l2 {
		for _, l1v := range l1 {
			tempR = append(tempR, l1v+l2v)
		}
	}
	dl[t] = tempR
	helper(dl, t, t+1)
}
