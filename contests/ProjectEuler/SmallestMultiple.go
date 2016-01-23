/*
This problem is a programming version of Problem 5 from projecteuler.net

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible(divisible with no remainder) by all of the numbers from 1 to N?

Input Format
First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

Output Format
Print the required answer for each test case.

Constraints
1≤T≤10
1≤N≤40

Sample Input

2
3
10
Sample Output

6
2520
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, res int
	var found bool
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		res = 1
		found = false
		for !found {
			found = true
			for j := 1; j <= n; j++ {
				if (res % j) != 0 {
					found = false
					break
				}
			}
			if found {
				fmt.Println(res)
			} else {
				res++
			}
		}
	}
}
