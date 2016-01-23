/*
This problem is a programming version of Problem 13 from projecteuler.net

Work out the first ten digits of the sum of N 50-digit numbers.

Input Format
First line contains N, next N lines contain a 50 digit number each.

Output Format
Print only first 10 digit of the final sum

Constraints
1≤N≤103

Sample Input

5
37107287533902102798797998220837590246510135740250
46376937677490009712648124896970078050417018260538
74324986199524741059474233309513058123726617309629
91942213363574161572522430563301811072406154908250
23067588207539346171171980310421047513778063246676
Sample Output

2728190129
*/

package main

import "fmt"

func array2int(a []byte) int {
	t, res := 1, 0
	for i, _ := range a {
		res += int(a[i]) * t
		t *= 10
	}
	return res
}

func arrayAdd(a, b []byte) []byte {
	var min, max []byte
	if len(a) < len(b) {
		min = a[:]
		max = b[:]
	} else {
		min = b[:]
		max = a[:]
	}
	var n, over byte
	res := make([]byte, 0)
	for i := 0; i < len(max); i++ {
		if i < len(min) {
			n = min[i] + max[i] + over
		} else {
			n = max[i] + over
		}
		res = append(res, n%10)
		over = n / 10
	}
	if over > 0 {
		res = append(res, over)
	}
	return res
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	sum := make([]byte, 0)
	digits := make([]byte, 50)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &digits)
		for j := 0; j < 50; j++ {
			digits[j] -= '0'
		}
		for j := 0; j < 25; j++ {
			digits[j], digits[50-1-j] = digits[50-1-j], digits[j]
		}
		sum = arrayAdd(digits, sum)
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", sum[len(sum)-1-i])
	}
	fmt.Println()
}
