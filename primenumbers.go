package main

import "fmt"

func main() {
	fmt.Println("Test if number is Prime")
	fmt.Println("4 = ", isPrime(4))
	fmt.Println("5 = ", isPrime(5))

	fmt.Println("\nPrint prime numbers till N")
	fmt.Println("primes till 11 ")
	fmt.Println(printPrimes(11))
}

//check if number is prime
func isPrime(N int) bool {
	// 1 is NOT prime
	if N <= 1 {
		return false
	}

	for i := 2; i < N; i++ {
		// fmt.Printf("checking if %d can be divided by %d \n", N, i)
		//can N be divided
		if N%i == 0 {
			return false
		}
	}
	return true
}

func printPrimes(N int) []int {
	primes := []int{}
	for i := 2; i <= N; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
