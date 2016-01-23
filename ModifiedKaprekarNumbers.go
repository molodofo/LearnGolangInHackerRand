/*
A modified Kaprekar number is a positive whole number n with d digits, such that when we split its square into two pieces - a right hand piece r with d digits and a left hand piece l that contains the remaining d or d−1 digits, the sum of the pieces is equal to the original number (i.e. l + r = n).

Note: r may have leading zeros.

Here's an explanation from Wikipedia about the ORIGINAL Kaprekar Number (spot the difference!): In mathematics, a Kaprekar number for a given base is a non-negative integer, the representation of whose square in that base can be split into two parts that add up to the original number again. For instance, 45 is a Kaprekar number, because 45² = 2025 and 20+25 = 45.

The Task
You are given the two positive integers p and q, where p is lower than q. Write a program to determine how many Kaprekar numbers are there in the range between p and q (both inclusive) and display them all.

Input Format

There will be two lines of input: p, lowest value q, highest value

Constraints:
0<p<q<100000
Output Format

Output each Kaprekar number in the given range, space-separated on a single line. If no Kaprekar numbers exist in the given range, print INVALID RANGE.

Sample Input

1
100
Sample Output

1 9 45 55 99

Explanation

1, 9, 45, 55, and 99 are the Kaprekar Numbers in the given range.
*/

package main

import "fmt"

func overturnByteArray(a []byte) []byte {
	for i := 0; i < (len(a) / 2); i++ {
		a[i], a[len(a)-i-1] = a[len(a)-i-1], a[i]
	}
	return a
}

func uint64ToByteArray(n uint64) []byte {
	a := make([]byte, 0)
	for i := 0; n > 0; i++ {
		a = append(a, byte(n%10))
		n /= 10
	}
	overturnByteArray(a)
	return a
}

func byteArrayToUint64(a []byte) (n uint64) {
	m := uint64(1)
	for i := len(a) - 1; i >= 0; i-- {
		n += uint64(a[i]) * m
		m *= 10
	}
	return
}

func splitTheNum(n uint64) (l, r uint64) {
	if n < 10 {
		l, r = 0, n
	} else {
		numArray := uint64ToByteArray(n)
		l, r = byteArrayToUint64(numArray[:len(numArray)/2]), byteArrayToUint64(numArray[len(numArray)/2:])
	}
	return
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var p, q, square, l, r uint64
	kn := make([]uint64, 0)
	fmt.Scanf("%d", &p)
	fmt.Scanf("%d", &q)
	//p, q = 1, 100
	for i := p; i <= q; i++ {
		square = i * i
		l, r = splitTheNum(square)
		if i == (l + r) {
			kn = append(kn, i)
		}
	}
	for i := 0; i < (len(kn) - 1); i++ {
		fmt.Printf("%d ", kn[i])
	}
	if len(kn) > 0 {
		fmt.Println(kn[len(kn)-1])
	} else {
		fmt.Println("INVALID RANGE")
	}
}
