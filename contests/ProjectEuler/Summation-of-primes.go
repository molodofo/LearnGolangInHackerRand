/*
This problem is a programming version of Problem 10 from projecteuler.net

The sum of the primes below 10 is 2+3+5+7=17.

Find the sum of all the primes not greater than given N.

Input Format
The first line contains an integer T i.e. number of the test cases.
The next T lines will contains an integer N.

Output Format
Print the value corresponding to each test case in seperate line.

Constraints
1≤T≤104
1≤N≤106

Sample Input

2
5
10
Sample Output

10
17
*/

package main

import "fmt"
import "math"

func isPrime(n int) bool {
	s := int(math.Sqrt(float64(n)))
	for i := 2; i <= s; i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, sum int
	fmt.Scanf("%d", &t)
	prime := make([]int, 1)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		sum = 0
		if n > prime[len(prime)-1] {
			for j := prime[len(prime)-1] + 1; j <= n; j++ {
				if isPrime(j) {
					prime = append(prime, j)
				}
			}
		}
		for j := 0; j < len(prime); j++ {
			if prime[j] > n {
				break
			}
			sum += prime[j]
		}
		fmt.Println(sum - 1)
	}
}
