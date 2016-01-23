/*
Given a time in AM/PM format, convert it to military (24-hour) time.

Note: Midnight is 12:00:00AM on a 12-hour clock and 00:00:00 on a 24-hour clock. Noon is 12:00:00PM on a 12-hour clock and 12:00:00 on a 24-hour clock.

Input Format

A time in 12-hour clock format (i.e.: hh:mm:ssAM or hh:mm:ssPM), where 01≤hh≤12.

Output Format

Convert and print the given time in 24-hour format, where 00≤hh≤23.

Sample Input

07:05:45PM
Sample Output

19:05:45
Explanation

7 PM is after noon, so you need to add 12 hours to it during conversion. 12 + 7 = 19. Minutes and seconds do not change in 12-24 hour time conversions, so the answer is 19:05:45.
*/

package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var time [10]int
	var hour int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%c", &time[i])
	}
	hour = ((time[0] - '0') * 10) + (time[1] - '0')
	switch time[8] {
	case 'P':
		if hour != 12 {
			hour += 12
			time[0] = (hour / 10) + '0'
			time[1] = (hour % 10) + '0'
		}
	case 'A':
		if hour == 12 {
			time[0] = '0'
			time[1] = '0'
		}
	}

	for i := 0; i < 8; i++ {
		fmt.Printf("%s", string(time[i]))
	}
	fmt.Println()
}
