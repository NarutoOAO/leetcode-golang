package leetcode

const mx = 1e6

var primes []int
var isPrimes = [mx + 1]bool{}

func init() {
	for i := 0; i < mx+1; i++ {
		isPrimes[i] = true
	}
	for i := 2; i < mx+1; i++ {
		if isPrimes[i] {
			primes = append(primes, i)
			for j := i * i; j < mx+1; j += i {
				isPrimes[j] = false
			}
		}
	}
}
