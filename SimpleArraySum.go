/*
You are given an array of integers of size N. Can you find the sum of the elements in the array?

Input 
The first line of input consists of an integer N. The next line contains N space-separated integers representing the array elements. 
Sample:
6
1 2 3 4 10 11

Output 
Output a single value equal to the sum of the elements in the array. 
For the sample above you would just print 31 since 1+2+3+4+10+11=31.
*/

package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    len_arr := 0
    fmt.Scanf("%d", &len_arr)
    //arr := make([]int, len_arr)
    sum := 0
    n := 0
    for i:=0; i<len_arr; i++ {
        fmt.Scanf("%d", &n)
        sum += n
    }
    fmt.Println(sum)
}
