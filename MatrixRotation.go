/*
You are given a 2D matrix, a, of dimension MxN and a positive integer R. You have to rotate the matrix R times and print the resultant matrix. Rotation should be in anti-clockwise direction.

Rotation of a 4x5 matrix is represented by the following figure. Note that in one rotation, you have to shift elements by one step only (refer sample tests for more clarity).

matrix-rotation

It is guaranteed that the minimum of M and N will be even.

Input Format
First line contains three space separated integers, M, N and R, where M is the number of rows, N is number of columns in matrix, and R is the number of times the matrix has to be rotated.
Then M lines follow, where each line contains N space separated positive integers. These M lines represent the matrix.

Output Format
Print the rotated matrix.

Constraints
2 <= M, N <= 300
1 <= R <= 109
min(M, N) % 2 == 0
1 <= aij <= 108, where i ∈ [1..M] & j ∈ [1..N]

Sample Input #00

4 4 1
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
Sample Output #00

2 3 4 8
1 7 11 12
5 6 10 16
9 13 14 15
Sample Input #01

4 4 2
1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16
Sample Output #01

3 4 8 12
2 11 10 16
1 7 6 15
5 9 13 14
Sample Input #02

5 4 7
1 2 3 4
7 8 9 10
13 14 15 16
19 20 21 22
25 26 27 28
Sample Output #02

28 27 26 25
22 9 15 19
16 8 21 13
10 14 20 7
4 3 2 1
Sample Input #03

2 2 3
1 1
1 1
Sample Output #03

1 1
1 1
Explanation
Sample Case #00: Here is an illustration of what happens when the matrix is rotated once.

 1  2  3  4      2  3  4  8
 5  6  7  8      1  7 11 12
 9 10 11 12  ->  5  6 10 16
13 14 15 16      9 13 14 15
Sample Case #01: Here is what happens when to the matrix after two rotations.

 1  2  3  4      2  3  4  8      3  4  8 12
 5  6  7  8      1  7 11 12      2 11 10 16
 9 10 11 12  ->  5  6 10 16  ->  1  7  6 15
13 14 15 16      9 13 14 15      5  9 13 14
Sample Case #02: Following are the intermediate states.

1  2  3  4      2  3  4 10    3  4 10 16    4 10 16 22
7  8  9 10      1  9 15 16    2 15 21 22    3 21 20 28
13 14 15 16 ->  7  8 21 22 -> 1  9 20 28 -> 2 15 14 27 ->
19 20 21 22    13 14 20 28    7  8 14 27    1  9  8 26
25 26 27 28    19 25 26 27    13 19 25 26   7 13 19 25



10 16 22 28    16 22 28 27    22 28 27 26    28 27 26 25
 4 20 14 27    10 14  8 26    16  8  9 25    22  9 15 19
 3 21  8 26 ->  4 20  9 25 -> 10 14 15 19 -> 16  8 21 13
 2 15  9 25     3 21 15 19     4 20 21 13    10 14 20  7
 1  7 13 19     2  1  7 13     3  2  1  7     4  3  2  1
Sample Case #03: As all elements are same, any rotation will reflect the same matrix.
*/

package main

import "fmt"

func min(m, n uint64) uint64 {
	if m < n {
		return m
	} else {
		return n
	}
}

func rotate(m, n, r, x1, y1 uint64) (x2, y2 uint64) {
	x0 := min(x1, m-x1-1)
	y0 := min(y1, n-y1-1)
	x0, y0 = min(x0, y0), min(x0, y0)
	m1, n1 := m-(x0*2), n-(y0*2)
	if (x1 >= x0) && (y1 == y0) {
		//fmt.Printf("b1: ")
		r += (x1 - x0)
	} else if x1 == (x0 + m1 - 1) {
		//fmt.Printf("b21: ")
		r += (m1 - 1 + (y1 - y0))
	} else if y1 == (y0 + n1 - 1) {
		//fmt.Printf("b3: ")
		r += (m1 + n1 - 1 + (m1 - (x1 - x0) - 2))
	} else {
		//fmt.Printf("b4: ")
		r += ((m1+n1)*2 - 4 - (y1 - y0))
	}
	r = r % ((m1+n1)*2 - 4)
	switch {
	case r < m1:
		//fmt.Printf("c1: ")
		x2, y2 = (x0 + r), y0
	case r < (m1 + n1 - 1):
		//fmt.Printf("c2: ")
		x2, y2 = (x0 + m1 - 1), (y0 + (r - m1 + 1))
	case r < (m1 + n1 - 1 + m1 - 1):
		//fmt.Printf("c3: ")
		x2, y2 = (x0 + (m1 - (r - (m1 + n1 - 1)) - 2)), (y0 + n1 - 1)
	default:
		//fmt.Printf("c4: ")
		x2, y2 = x0, (y0 + (((m1+n1)*2 - 4) - r))
	}
	//fmt.Println("r =", r, ",", x0, y0, ",", x1, y1, ",", x2, y2, ",", m1, n1)
	return
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	var m, n, r uint64
	fmt.Scanf("%d %d %d", &m, &n, &r)

	matrix := make([][]uint64, m)
	matrix2 := make([][]uint64, m)
	for i := uint64(0); i < m; i++ {
		matrix[i] = make([]uint64, n)
		matrix2[i] = make([]uint64, n)
	}
	for i := uint64(0); i < m; i++ {
		for j := uint64(0); j < n; j++ {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	for i, i1, j1 := uint64(0), uint64(0), uint64(0); i < m; i++ {
		for j := uint64(0); j < n; j++ {
			i1, j1 = rotate(m, n, r, i, j)
			matrix2[i1][j1] = matrix[i][j]
		}
	}

	for i := uint64(0); i < m; i++ {
		for j := uint64(0); j < (n - 1); j++ {
			fmt.Printf("%d ", matrix2[i][j])
		}
		fmt.Println(matrix2[i][n-1])
	}
}
