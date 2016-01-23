/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of a given number N?

Input Format
First line contains T, the number of test cases. This is followed by T lines each containing an integer N.

Output Format
For each test case, display the largest prime factor of N.

Constraints
1≤T≤10
10≤N≤1012
Sample Input

2
10
17
Sample Output

5
17
*/

package main

import "fmt"
import "math"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, s int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		s = int(math.Sqrt(float64(n)))
		for j := 2; j <= s; j++ {
			for (n % j) == 0 {
				n /= j
			}
			if n == 1 {
				n = j
				break
			}
		}
		fmt.Println(n)
	}
}
