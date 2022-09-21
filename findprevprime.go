package main

// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if isPrime1(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}

func isPrime1(num int) bool {
	if num == 1 { // 1 is not prime
		return false
	}
	for i := 2; i <= num/2; i++ { // only go up to the num's half
		if num%i == 0 {
			return false
		}
	}
	return true // only got here if the loop didn't return false, which means that the number was never divided perfectly by any number inferior to its half
}

// EXTRA TIP
// 2 is the smallest prime number. It also the only even prime number.

// TESTING BELOW

// func main() {
// 	fmt.Println(FindPrevPrime(5)) // 5
// 	fmt.Println(FindPrevPrime(4)) // 3
// 	fmt.Println(FindPrevPrime(1)) // 0
// 	fmt.Println(FindPrevPrime(0)) // 0
// }
