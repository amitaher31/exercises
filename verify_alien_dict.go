//https://leetcode.com/problems/verifying-an-alien-dictionary/solution/
package main

import "fmt"

func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(words, order, isAlienSorted(words, order))
}

func isAlienSorted(words []string, order string) bool {
	if len(words) < 2 {
		return true
	}

	ordermap := make(map[uint8]int, len(order))
	for i, _ := range order {
		ordermap[order[i]] = i
	}
	fmt.Println(ordermap)

	for i, j := 0, 1; j < len(words); i, j = i+1, j+1 {
		a, b := words[i], words[j]
		fmt.Println(a, b)

		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] == b[k] {
				continue
			}

			if ordermap[a[k]] > ordermap[b[k]] || len(a) > len(b) {
				return false
			}
		}
	}

	return true
}
