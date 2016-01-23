/*
Your teacher has given you the task of drawing a staircase structure. Being an expert programmer, you decided to make a program to draw it for you instead. Given the required height, can you print a staircase as shown in the example?

Input
You are given an integer N depicting the height of the staircase.

Output
Print a staircase of height N that consists of # symbols and spaces. For example for N=6, here's a staircase of that height:

     #
    ##
   ###
  ####
 #####
######
Note: The last line has 0 spaces before it.
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	n := 0
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= (n - i); j++ {
			fmt.Printf(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}
}
