/*

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

*/

package main

import (
	"fmt"
)

func main() {

	limit := 20
	ans := solve(limit)
	fmt.Println(ans)
}

func solve(limit int) int {

	product := 1
	for factor := 2; factor <= limit; factor++ {
		for smallerFactor := 1; smallerFactor <= factor; smallerFactor++ {
			if smallerFactor*product%factor == 0 {
				product *= smallerFactor
				break
			}
		}
	}
	return product
}
