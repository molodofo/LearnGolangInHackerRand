/*
This problem is a programming version of Problem 4 from projecteuler.net

A palindromic number reads the same both ways. The smallest 6 digit palindrome made from the product of two 3-digit numbers is 101101=143×707.

Find the largest palindrome made from the product of two 3-digit numbers which is less than N.

Input Format
First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

Output Format
Print the required answer for each test case in a new line.

Constraints
1≤T≤100
101101<N<1000000
Sample Input

2
101110
800000
Sample Output

101101
793397
*/

package main

import "fmt"

func int2array(n int) []byte {
	a := make([]byte, 0)
	for n > 0 {
		a = append(a, byte(n%10))
		n = n / 10
	}
	return a
}

func isPalindrome(n int) bool {
	a := int2array(n)
	for i := 0; i < (len(a) / 2); i++ {
		if a[i] != a[len(a)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n int
	var found bool
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		found = false
		for j := n; j > 10000; j-- {
			if isPalindrome(j) {
				for k := 100; k < 1000; k++ {
					if (j%k) == 0 && (j/k) < 1000 {
						found = true
						fmt.Println(j)
						break
					}
				}
				if found {
					break
				}
			}
		}
	}
}
