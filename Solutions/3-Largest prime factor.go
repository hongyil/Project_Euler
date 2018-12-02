/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

*/

package main

import (
	"fmt"
)

func main() {

	num := 600851475143
	ans := solve(num)
	fmt.Println(ans)
}

func solve(num int) int {

	factor := 2
	for num != 1 {
		if valid(factor) {
			for num%factor == 0 {
				num /= factor
			}
		}
		factor += 1
	}

	return factor - 1
}

func valid(num int) bool {

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}
