/*
Taum正在计划为他的朋友Diksha庆祝生日。Diksha想找Taum要两种礼物：一种是黑的，一种是白的。 为了让她高兴， Taum要买B件黑色的礼物和W件白色的礼物。

每件黑色的礼物花费为X个单位
每件白色的礼物花费为Y个单位
把一件黑色的礼物转变为白色的，或者把一件白色的礼物转变为黑色的花费是Z个单位。
帮助Taum计算他要给Diksha买礼物的最小花费。

Input Format

第一行包含一个整数T，表示测试数据的组数。
后面有T个双行。每组测试数据第一行包含B和W的值，每组数据的另外一行包含X,Y和Z的值。

约束条件
1≤T≤10
0≤X,Y,Z,B,W≤109

Output Format

T行，每行包含对应测试数据的答案。

Sample Input

1
10 10
1 1 1
Sample Output

20

Explanation

把白色转变为黑色或者黑色转变为白色没有任何好处，所以他需要用1个单位的花费买每个礼物。
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var t, x, y, z, b, w uint
	fmt.Scanf("%d", &t)
	for i := uint(0); i < t; i++ {
		fmt.Scanf("%d %d", &b, &w)
		fmt.Scanf("%d %d %d", &x, &y, &z)
		if x > (y + z) {
			fmt.Println((b * (y + z)) + (w * y))
		} else if y > (x + z) {
			fmt.Println((b * x) + (w * (x + z)))
		} else {
			fmt.Println((b * x) + (w * y))
		}
	}
}
