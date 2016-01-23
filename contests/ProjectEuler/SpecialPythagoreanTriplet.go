/*
This problem is a programming version of Problem 9 from projecteuler.net

A Pythagorean triplet is a set of three natural numbers, a<b<c, for which,
a2+b2=c2
For example, 3^2+4^2=9+16=25=5^2
Given N, Check if there exists any Pythagorean triplet for which a+b+c=N
Find maximum possible value of abc among all such Pythagorean triplets, If there is no such Pythagorean triplet print −1.

Input Format
The first line contains an integer T i.e. number of test cases.
The next T lines will contain an integer N.

Output Format
Print the value corresponding to each test case in separate lines.

Constraints
1≤T≤3000
1≤N≤3000
Sample Input

2
12
4
Sample Output

60
-1
*/

package main

import "fmt"

func isPythagoreanTriplet(a, b, c int) bool {
	if ((a * a) + (b * b)) == (c * c) {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, a, b, c, res int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		res = -1
		for a = 1; a < (n / 2); a++ {
			b = ((n * n) - (2 * a * n)) / ((2 * n) - (2 * a))
			c = n - a - b
			if isPythagoreanTriplet(a, b, c) {
				res = max(a*b*c, res)
			}
		}
		fmt.Println(res)
	}
}
