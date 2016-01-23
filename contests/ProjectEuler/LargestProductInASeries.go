/*
This problem is a programming version of Problem 8 from projecteuler.net

Find the greatest product of K consecutive digits in the N digit number.

Input Format
First line contains T that denotes the number of test cases.
First line of each test case will contain two integers N & K.
Second line of each test case will contain a N digit integer.

Output Format
Print the required answer for each test case.

Constraints
1≤T≤100
1≤K≤7
K≤N≤1000

Sample Input

2
10 5
3675356291
10 5
2709360626
Sample Output

3150
0
*/

package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, k, n, pro, res int
	s := make([]byte, 0)
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		s = make([]byte, n)
		fmt.Scanf("%s", &s)
		pro, res = 1, 0
		for j := 0; j < (n - k - 1); j++ {
			for x := 0; x < k; x++ {
				pro *= int(s[j+x] - '0')
			}
			res = max(pro, res)
			pro = 1
		}
		fmt.Println(res)
	}
}
