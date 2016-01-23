/*
This problem is a programming version of Problem 6 from projecteuler.net

The sum of the squares of the first ten natural numbers is, 12+22+...+102=385. The square of the sum of the first ten natural numbers is, (1+2+⋯+10)2=552=3025. Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025−385=2640.

Find the difference between the sum of the squares of the first N natural numbers and the square of the sum.

Input Format
First line contains T that denotes the number of test cases. This is followed by T lines, each containing an integer, N.

Output Format
Print the required answer for each test case.

Constraints
1≤T≤104
1≤N≤104

Sample Input

2
3
10
Sample Output

22
2640
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, n, square_sum, sum_square int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &n)
		square_sum = n * (n + 1) * ((2 * n) + 1) / 6
		sum_square = n * (n + 1) / 2
		sum_square *= sum_square
		fmt.Println(sum_square - square_sum)
	}
}
