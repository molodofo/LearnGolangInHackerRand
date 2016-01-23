/*
Given a square matrix of size N×N, calculate the absolute difference between the sums of its diagonals.

Input Format

The first line contains a single integer, N. The next N lines denote the matrix's rows, with each line containing N space-separated integers describing the columns.

Output Format

Print the absolute difference between the two sums of the matrix's diagonals as a single integer.

Sample Input
3
11 2 4
4 5 6
10 8 -12

Sample Output
15

Explanation
The primary diagonal is:
11
      5
            -12

Sum across the primary diagonal: 11 + 5 - 12 = 4

The secondary diagonal is:
            4
      5
10
Sum across the secondary diagonal: 4 + 5 + 10 = 19
Difference: |4 - 19| = 15
*/
package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	n := 0
	fmt.Scanf("%d", &n)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	diff := 0
	for i := 0; i < n; i++ {
		diff += matrix[i][i]
	}
	for i := 0; i < n; i++ {
		diff -= matrix[i][n-1-i]
	}
	if diff < 0 {
		diff = -diff
	}
	fmt.Println(diff)
}
