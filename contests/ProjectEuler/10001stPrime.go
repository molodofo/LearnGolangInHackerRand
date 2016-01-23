/*
This problem is a programming version of Problem 7 from projecteuler.net

By listing the first six prime numbers: 2,3,5,7,11 and 13, we can see that the 6th prime is 13.
What is the N'th prime number?

Input Format
First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

Output Format
Print the required answer for each test case.

Constraints
1≤T≤103
1≤N≤104

Sample Input

2
3
6
Sample Output

5
13
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
	var t, n, res int
	prime := make([]int, 0)
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		if n > len(prime) {
			if len(prime) > 0 {
				res = prime[len(prime)-1]
			} else {
				res = 1
			}
			for j := len(prime); j < n; j++ {
				for true {
					res++
					if isPrime(res) {
						prime = append(prime, res)
						break
					}
				}
			}
		}
		fmt.Println(prime[n-1])
	}
}
