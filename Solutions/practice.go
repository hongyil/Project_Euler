/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

*/

package main

import (
	"fmt"
)

var data [][]int

func main() {

	for a := 2; a <= 100; a++ {
		for b := 2; b <= 100; b++ {
			tmp := pow(a, b)
			if !contain(data, tmp) {
				data = append(data, tmp)
			}
		}
	}
	fmt.Println(len(data))

}

func pow(a, n int) []int {

	digits := make([]int, 10001)
	digits[0] = a
	num := 1
	for i := 1; i < n; i++ {
		carry := 0
		for j := 0; j < num; j++ {
			tmp := digits[j]*a + carry
			digits[j] = tmp % 10
			carry = tmp / 10
		}
		for carry > 0 {
			digits[num] = carry % 10
			carry /= 10
			num += 1
		}
	}
	return digits
}

func contain(data [][]int, b []int) bool {

	if len(data) == 0 {
		return false
	}
	for _, a := range data {
		if equals(a, b) {
			return true
		}
	}
	return false
}

func equals(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
