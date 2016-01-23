/*
Given the time in numerals we may convert it into words, as shown below:

5:00→ five o' clock
5:01→ one minute past five
5:10→ ten minutes past five
5:30→ half past five
5:40→ twenty minutes to six
5:45→ quarter to six
5:47→ thirteen minutes to six
5:28→ twenty eight minutes past five
Write a program which prints the time in words for the input given in the format mentioned above.

Input Format

There will be two lines of input:
H, representing the hours
M, representing the minutes

Constraints
1≤H<12
0≤M<60
Output Format

Display the time in words.

Sample Input

5
47
Sample Output

thirteen minutes to six
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	numStr := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
	}
	var hour, min int
	fmt.Scanf("%d", &hour)
	fmt.Scanf("%d", &min)
	if min == 0 {
		fmt.Printf(numStr[hour] + " o' clock\n")
	} else if min <= 20 {
		if min == 15 {
			fmt.Printf("quarter past " + numStr[hour])
		} else {
			fmt.Printf(numStr[min])
			if min == 1 {
				fmt.Printf(" minute past ")
			} else {
				fmt.Printf(" minutes past ")
			}
			fmt.Printf(numStr[hour] + "\n")
		}
	} else if min < 30 {
		fmt.Printf(numStr[20] + " " + numStr[min%10] + " minutes past " + numStr[hour] + "\n")
	} else if min > 30 {
		if min < 40 {
			fmt.Printf(numStr[20] + " " + numStr[(60-min)%10] + " minutes to " + numStr[hour+1] + "\n")
		} else {
			if min == 45 {
				fmt.Printf("quarter to " + numStr[hour+1])
			} else {
				fmt.Printf(numStr[60-min] + " minutes to " + numStr[hour+1] + "\n")
			}
		}
	} else if min == 30 {
		fmt.Printf("half past " + numStr[hour] + "\n")
	}
}
