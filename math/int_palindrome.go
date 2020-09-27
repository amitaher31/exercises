/* 9. Palindrome Number
https://leetcode.com/problems/palindrome-number/

w/o converting to string & using leetcodes algo
 - reverse the number & compare with original num, if they are equal => palindrome
 - The reverted num can be larger than int.Max
  - so we can reverse right half of the num & compare with left half to determine if its palindrome
  - Negative number cannot be a palindrome (since there isn't a `-` sign on right ever)

For example, if the input is 1221, if we can revert the last part of the number "1221" from "21" to "12", and compare it with the first half of the number "12", since 12 is the same as 12, we know that the number is a palindrome.

Let's see how we could translate this idea into an algorithm.

Algorithm

First of all we should take care of some edge cases. All negative numbers are not palindrome, for example: -123 is not a palindrome since the '-' does not equal to '3'. So we can return false for all negative numbers.

Now let's think about how to revert the last half of the number. For number 1221, if we do 1221 % 10, we get the last digit 1, to get the second to the last digit, we need to remove the last digit from 1221, we could do so by dividing it by 10, 1221 / 10 = 122. Then we can get the last digit again by doing a modulus by 10, 122 % 10 = 2, and if we multiply the last digit by 10 and add the second last digit, 1 * 10 + 2 = 12, it gives us the reverted number we want. Continuing this process would give us the reverted number with more digits.

Now the question is, how do we know that we've reached the half of the number?

Since we divided the number by 10, and multiplied the reversed number by 10, when the original number is less than the reversed number, it means we've processed half of the number digits.
*/

package main

import "fmt"

func main() {
	fmt.Println("121 is palindrome = ", isPalindrome(121))
	fmt.Println("-121 is palindrome = ", isPalindrome(-121))
	fmt.Println("0 is palindrome = ", isPalindrome(0))
	fmt.Println("10 is palindrome = ", isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rightReverse := 0
	for x > rightReverse {
		// get last digit
		last := x % 10
		// 2ndLast := x / 10
		// x /= 10
		x = x / 10
		// fmt.Println(x)
		rightReverse = rightReverse*10 + last
		fmt.Println(x, rightReverse)
	}

	return x == rightReverse || x == rightReverse/10
}

// import "strconv"
// func isPalindrome(x int) bool {
//     xStr := strconv.Itoa(x)
//     // fmt.Println(xStr)
//     l, r := 0, len(xStr)-1
//     // fmt.Println(l,r)

//     for l < r {
//         if xStr[l] != xStr[r] {
//             return false
//         }
//         l++
//         r--
//     }

//     return true
// }
