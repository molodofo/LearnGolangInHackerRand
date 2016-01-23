/*
You are given an integer N. Print the factorial of this number.

N!=N×(N−1)×(N−2)×⋯×3×2×1
Input
Input consists of a single integer N, where 1≤N≤100.

Output
Print the factorial of N.

Example
For an input of 25, you would print 15511210043330985984000000.

Note: Factorials of N>20 can't be stored even in a 64−bit long long variable. Big integers must be used for such calculations. Languages like Java, Python, Ruby etc. can handle big integers, but we need to write additional code in C/C++ to handle huge values.

We recommend solving this challenge using BigIntegers.
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

func add(a, b []byte) []byte {
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

// multiplication
func multiply(a, b []byte) []byte {
	c := make([]byte, 0)
	res := make([]byte, 0)
	var n, over byte
	for i := 0; i < len(b); i++ {
		c = make([]byte, i)
		for j := 0; j < len(a); j++ {
			n = (a[j] * b[i]) + over
			c = append(c, n%10)
			over = n / 10
		}
		if over > 0 {
			c = append(c, over)
			over = 0
		}
		res = add(res, c)
	}
	return res
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scanf("%d", &n)
	res := make([]byte, 0)
	res = append(res, byte(1))
	for i := 2; i <= n; i++ {
		res = multiply(res, int2array(i))
	}
	for i := (len(res) - 1); i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
	fmt.Println()
}
