package main

import "fmt"

func main() {
	fmt.Println("USA: ", detectCapitalUse("USA"))
	fmt.Println("USa: ", detectCapitalUse("USa"))
	fmt.Println("usa: ", detectCapitalUse("usa"))
	fmt.Println("usA: ", detectCapitalUse("usA"))
	fmt.Println("Usa: ", detectCapitalUse("Usa"))
}

func detectCapitalUse(word string) bool {
	// fmt.Println('A', 'a', word[0])

	if len(word) <= 1 {
		return true
	}

	startCaps := word[0] < 'a'

	hasSmall, hasUpper := false, false
	if word[1] < 'a' {
		hasUpper = true
	} else {
		hasSmall = true
	}

	for i := 1; i < len(word); i++ {
		// fmt.Println(word[i])

		//uSa
		if !startCaps && word[i] < 'a' {
			return false
		}

		//UsA
		if startCaps && hasSmall && word[i] < 'a' {
			return false
		}

		//USa
		if startCaps && hasUpper && word[i] >= 'a' {
			return false
		}

	}
	return true
}
