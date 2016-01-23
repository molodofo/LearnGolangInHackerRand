/*
This problem is a programming version of Problem 1 from projecteuler.net

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below N.

Input Format
First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

Output Format
For each test case, print an integer that denotes the sum of all the multiples of 3 or 5 below N.

Constraints
1≤T≤105
1≤N≤109

Sample Input

2
10
100
Sample Output

23
2318
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, sum, three, five, tf int
	//var sum int
	fmt.Scanf("%d", &t)
	for i := int(0); i < t; i++ {
		fmt.Scanf("%d", &n)
		sum = int(0)
		for j := (n - 1); j > 0; j-- {
			if (j % int(3)) == 0 {
				three = j
				break
			}
		}
		sum += ((three / 3) * (three + 3)) / 2
		for j := (n - 1); j > 0; j-- {
			if (j % int(5)) == 0 {
				five = j
				break
			}
		}
		sum += ((five / 5) * (five + 5)) / 2
		for j := (n - 1); j > 0; j-- {
			if (j % int(15)) == 0 {
				tf = j
				break
			}
		}
		sum -= ((tf / 15) * (tf + 15)) / 2
		fmt.Println(sum)
	}
}
