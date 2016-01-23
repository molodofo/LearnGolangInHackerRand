/*
给你一个n×n的地图。地图中的每个格子有一个值表示该地区的深度。我们称一个地图中的一个格子为空洞，当且仅当该格子不在地图边缘并且每个和它相邻的格子都具有比它更小的深度。两个格子称为相邻如果它们共有一条边。

你要找到地图中所有的空洞，并且用X描述。

输入格式

第一行包含一个整数n，表示地图的规模。 接下来n行中每行包含n个无空白的正数字。每个数字（1-9）表示对应区域的深度。

输出格式

输出n行，表示最终的地图结果。每个空洞要用字符X替换。

约束条件

1≤n≤100
输入样例

4
1112
1912
1892
1234
输出样例

1112
1X12
18X2
1234
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n, enter int
	fmt.Scanf("%d", &n)
	m := make([][]int, n)
	m2 := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		m2[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scanf("%c", &m[i][j])
			m2[i][j] = m[i][j]
		}
		fmt.Scanf("%c", &enter) // enter
	}
	for i := 1; i < (n - 1); i++ {
		for j := 1; j < (n - 1); j++ {
			if m[i][j] > m[i-1][j] &&
				m[i][j] > m[i+1][j] &&
				m[i][j] > m[i][j-1] &&
				m[i][j] > m[i][j+1] {
				m2[i][j] = 'X'
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%c", m2[i][j])
		}
		fmt.Printf("\n") // enter
	}
}
