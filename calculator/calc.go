package calc

import "fmt"

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6
func sumOfFirst(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + n - i
	}
	return sum
}

// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	count := 0
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			count++
			break
		}
	}
	if count == 0 && n != 1 {
		return true
	} else {
		return false
	}
}

func main() {
	sum := sumOfFirst(3)
	fmt.Println(sum)
}
